#!/usr/bin/env bash
# Phase019-A dogfood rehearsal — one full operating-loop revolution on
# fixtures, machine-judged end to end:
#
#   sync → ingest (fixture CF logs + GSC stub) → freshness scan → queue
#   export → refresh-quest consumption (deliberate queue-scope FAIL, then
#   PASS) → ① article commit → re-sync (lastmod recognized) → /hooks/deployed
#   → /archive/process (stub receipts done) → ② queue-deletion commit →
#   export consumed sync → monthly report generated AND published to the
#   bare origin.
#
# Topology (hurl-test/queue-export-test precedents, all throwaway):
#   - postgres: initdb cluster on 127.0.0.1:55432, trust auth
#   - abloqd:   built from backend/arts/backend, listens :8080
#   - stubs:    backend/scripts/archive-stub (wayback/indexnow/GSC)
#   - blog:     a git instance whose baseURL/section/slugs join the
#               backend/fixtures/cflogs lines and the GSC stub rows
#               (https://fixture.example.com /tech/post-a|post-b)
#   - bare origin = the repo abloqd exports queues to, the agent clones
#     from, and the report publishes into; a separate "deploy" clone plays
#     the production checkout BLOG_REPO_PATH points at.
#
# Evidence (plan Phase019 §완료판정 A) is preserved in the record dir:
#   ① refresh-session.json        gate session (try1 FAIL → try2 PASS)
#   ② receipts-done.json          GET /receipts?...&status=done (3 rows)
#   ③ report-publish.txt          bare-origin commit hash of reports/<ym>.md
#   ④ export-consumed.json        POST /queue/export consumed sync + rows
# plus the issued queue file, the next-prompt, the FAIL feedback, the agent
# commit log and step-by-step API responses (steps.log).
#
# "No human in the loop": every verdict here is computed (gate rules, HTTP
# asserts, git asserts). The article edit itself is a scripted agent action.
#
# Phase020: the rehearsal keeps the BLOG_REPO_PATH-only boot — abloqd
# synthesizes the single `default` site at startup (backward compat) and
# every domain call rides under /sites/default/… (v0.2.0 path move).
#
# Usage: backend/scripts/rehearsal/run.sh [record-dir]
#   record-dir  default docs/rehearsal/2026-06-loop1
# Requires: backend/arts generated, /tmp/abloq-goproxy (local-goproxy.sh),
#           postgres 16 binaries, python3, openssl, git.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../../.." && pwd)"
RECORD="${1:-$ROOT/docs/rehearsal/2026-06-loop1}"
PG="${PG_BIN:-/usr/lib/postgresql/16/bin}"
WORK="$(mktemp -d)"
PGPORT=55432
PORT_STUB=18097
HOST="http://127.0.0.1:8080"
TODAY="$(date -u +%F)"
YM="$(date -u +%Y-%m)"
PIDS=()

cleanup() {
  for pid in "${PIDS[@]:-}"; do kill "$pid" 2>/dev/null || true; done
  "$PG/pg_ctl" -D "$WORK/pg" stop -m immediate >/dev/null 2>&1 || true
  rm -rf "$WORK"
}
trap cleanup EXIT

fail() { echo "FAIL: $*" >&2; exit 1; }
jsonget() { python3 -c 'import json,sys; print(json.load(sys.stdin)'"$1"')'; }
step() { echo "== $*" | tee -a "$RECORD/steps.log"; }

mkdir -p "$RECORD"
: > "$RECORD/steps.log"

# module resolution: the local file GOPROXY serves abloq v0.0.12 + reins/toulmin
if [ ! -d /tmp/abloq-goproxy ]; then
  "$ROOT/backend/scripts/local-goproxy.sh" >/dev/null
fi
export GOPROXY="file:///tmp/abloq-goproxy,https://proxy.golang.org,direct"
export GONOSUMDB="github.com/park-jun-woo/abloq,github.com/park-jun-woo/reins,github.com/park-jun-woo/toulmin"
export GOFLAGS=-mod=mod

# ─────────────────────────────────────────────────────────────────────────
step "0. build abloqd + abloq CLI, temp postgres (127.0.0.1:$PGPORT trust), stubs"

(cd "$ROOT/backend/arts/backend" && go build -o "$WORK/abloqd" ./cmd)
(cd "$ROOT" && go build -o "$WORK/abloq" ./cmd/abloq)

"$PG/initdb" -D "$WORK/pg" -A trust -U postgres >/dev/null
"$PG/pg_ctl" -D "$WORK/pg" -o "-p $PGPORT -k $WORK/pg -c listen_addresses=127.0.0.1" \
  -l "$WORK/pg.log" start >/dev/null
"$PG/createdb" -h 127.0.0.1 -p "$PGPORT" -U postgres abloqd
for mig in "$ROOT"/backend/arts/db/migrations/*.up.sql; do
  "$PG/psql" -q -h 127.0.0.1 -p "$PGPORT" -U postgres -d abloqd -f "$mig" >/dev/null
done

# operator seed (bcrypt via the arts module — hurl-test precedent)
cat > "$WORK/hashgen.go" <<'EOF'
//go:build ignore

package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	h, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(h))
}
EOF
OP_HASH=$(cd "$ROOT/backend/arts/backend" && go run "$WORK/hashgen.go" abloq-rehearsal-operator)
"$PG/psql" -q -h 127.0.0.1 -p "$PGPORT" -U postgres -d abloqd <<SQL
INSERT INTO users (email, password_hash, role)
VALUES ('operator@rehearsal.test', '$OP_HASH', 'operator');
SQL

python3 "$ROOT/backend/scripts/archive-stub/archive_stub.py" "$PORT_STUB" &
PIDS+=($!)

# throwaway GSC service-account JSON (the stub token endpoint accepts any key)
openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out "$WORK/sa.pem" 2>/dev/null
python3 - "$WORK/sa.pem" "$WORK/sa.json" <<'EOF'
import json, sys
json.dump({"client_email": "stub@rehearsal-project.iam.gserviceaccount.com",
           "private_key": open(sys.argv[1]).read()}, open(sys.argv[2], "w"))
EOF

# ─────────────────────────────────────────────────────────────────────────
step "1. blog instance (git) → bare origin → deploy checkout"
# baseURL/section/slugs join backend/fixtures/cflogs and the GSC stub rows.
# post-a is stale by construction (fixed past lastmod, freshness_days 30);
# post-b's lastmod is generated as today so it never goes stale on reruns.

SEED="$WORK/seed"
mkdir -p "$SEED/content/ko/tech"
cat > "$SEED/blog.yaml" <<'EOF'
site:
  baseURL: https://fixture.example.com
  title: Rehearsal Blog
  author: Rehearsal Operator
  default_lang_in_subdir: false

languages: [ko]
sections: [tech]

structure:
  order: [body, sources]
  headings:
    sources: { ko: "출처" }

geo:
  freshness_days: 30

deploy:
  provider: s3-cloudfront
EOF
printf '.abloq/\n' > "$SEED/.gitignore"

cat > "$SEED/content/ko/tech/post-a.md" <<'EOF'
---
title: "Post A"
date: 2026-03-28
lastmod: 2026-04-01
tags: [rehearsal]
---

![main](cover.png)

*Image: by Rehearsal Operator*

This stale body sentence still describes the situation as of early 2025 in vendor terms.

Throughput grew 40% in 2025 per the vendor study. [Vendor study](https://example.org/spec)

## 출처

- [Vendor study](https://example.org/spec)
EOF

cat > "$SEED/content/ko/tech/post-b.md" <<EOF
---
title: "Post B"
date: 2026-06-01
lastmod: $TODAY
tags: [rehearsal]
---

![main](cover.png)

*Image: by Rehearsal Operator*

Fresh body that links to [post-a](/tech/post-a/) and stays inside the freshness window.

## 출처

- [Reference note](https://example.org/ref)
EOF

git init -q -b main "$SEED"
git -C "$SEED" add -A
git -C "$SEED" -c user.name=seed -c user.email=seed@rehearsal.test commit -q -m "seed: rehearsal instance"

git init -q --bare -b main "$WORK/bare"
git -C "$SEED" push -q "file://$WORK/bare" main
git clone -q "file://$WORK/bare" "$WORK/deploy"   # the "production checkout"

# ─────────────────────────────────────────────────────────────────────────
step "2. abloqd up (stub-wired)"

env DATABASE_URL="postgres://postgres@127.0.0.1:$PGPORT/abloqd?sslmode=disable" \
  JWT_SECRET="rehearsal-secret-rehearsal-secret-0123456789" \
  OPA_POLICY_PATH="$ROOT/backend/arts/backend/policy/authz.rego" \
  GIN_MODE=release LOG_LEVEL=error \
  WAYBACK_BASE_URL="http://127.0.0.1:$PORT_STUB" \
  WAYBACK_ACCESS_KEY=stub WAYBACK_SECRET_KEY=stub \
  INDEXNOW_ENDPOINT="http://127.0.0.1:$PORT_STUB/indexnow" \
  INDEXNOW_KEY="rehearsalkey" \
  GSC_API_BASE="http://127.0.0.1:$PORT_STUB" \
  GSC_TOKEN_URL="http://127.0.0.1:$PORT_STUB/token" \
  GSC_SEARCH_API_BASE="http://127.0.0.1:$PORT_STUB" \
  GSC_SA_JSON_PATH="$WORK/sa.json" \
  GSC_LOOKBACK_DAYS=3 \
  BLOG_REPO_PATH="$WORK/deploy" \
  QUEUE_EXPORT_REPO_URL="file://$WORK/bare" \
  QUEUE_EXPORT_WORKDIR="$WORK/export" \
  CF_LOG_SOURCE="$ROOT/backend/fixtures/cflogs" \
  "$WORK/abloqd" > "$WORK/abloqd.log" 2>&1 &
PIDS+=($!)
for _ in $(seq 1 100); do
  curl -sf "$HOST/health" >/dev/null 2>&1 && break
  sleep 0.2
done
curl -sf "$HOST/health" >/dev/null || { tail -20 "$WORK/abloqd.log" >&2; fail "abloqd did not come up"; }

# one login for the whole run (rate limit 5/min/IP, token TTL 15min)
TOKEN=$(curl -sf -X POST "$HOST/auth/login" -H 'Content-Type: application/json' \
  -d '{"email":"operator@rehearsal.test","password":"abloq-rehearsal-operator"}' \
  | jsonget '["access_token"]')
[ -n "$TOKEN" ] || fail "login failed"
AUTH="Authorization: Bearer $TOKEN"

api() { # $1=method $2=path [$3=json-body] — logs the response, prints it
  local out
  if [ $# -ge 3 ]; then
    out=$(curl -sf -m 600 -X "$1" "$HOST$2" -H "$AUTH" -H 'Content-Type: application/json' -d "$3")
  else
    out=$(curl -sf -m 600 -X "$1" "$HOST$2" -H "$AUTH")
  fi
  printf '%s %s -> %s\n' "$1" "$2" "$out" >> "$RECORD/steps.log"
  printf '%s' "$out"
}

# ─────────────────────────────────────────────────────────────────────────
step "3. sync + fixture ingest (CF logs, GSC) + freshness scan + export"

SYNCED=$(api POST /sites/default/sync | jsonget '["synced"]')
[ "$SYNCED" = "2" ] || fail "sync: synced=$SYNCED, want 2"

api POST /sites/default/ingest/crawl >/dev/null
HITS=$(api GET /sites/default/crawl-hits | jsonget '["hits"].__len__()')
[ "$HITS" -ge 1 ] || fail "ingest/crawl: no crawl_hits rows landed"

api POST /sites/default/ingest/gsc '{"inspect":false}' >/dev/null

DETECTED=$(api POST /sites/default/scans/freshness '{"ym":""}' | jsonget '["detected"]')
[ "$DETECTED" = "1" ] || fail "scans/freshness: detected=$DETECTED, want 1 (post-a only)"

EXPORT1=$(api POST /sites/default/queue/export)
[ "$(echo "$EXPORT1" | jsonget '["exported"]')" = "1" ] || fail "export: $EXPORT1, want exported=1"

# ─────────────────────────────────────────────────────────────────────────
step "4. agent leg — clone, quest scan/next, deliberate FAIL, then PASS"

AGENT="$WORK/agent"
git clone -q "file://$WORK/bare" "$AGENT"
QFILE="quests/queue/refresh-ko-tech-post-a.yaml"
[ -f "$AGENT/$QFILE" ] || fail "queue file $QFILE not issued"
grep -q '^key: ' "$AGENT/$QFILE" || fail "queue file has no gate-contract key: line"
cp "$AGENT/$QFILE" "$RECORD/queue-file.yaml"

# session/results live OUTSIDE the instance: an untracked file inside the
# clone is an out-of-scope working-tree change and queue-scope would
# (correctly) FAIL it — Phase018 trial memo, same class as .abloq/.
SESSION="$WORK/refresh-session.json"
QUEST() { (cd "$AGENT" && "$WORK/abloq" quest refresh "$@" --session "$SESSION" --out "$WORK/refresh-results.jsonl"); }

QUEST scan . >> "$RECORD/steps.log"
QUEST next > "$RECORD/next-prompt.txt"

# the refresh edit (the agent's authored work, scripted deterministically):
# advance lastmod to today, rewrite the stale sentence, replace the figure —
# same-source numeric replacement, >= min_meaningful_diff tokens.
python3 - "$AGENT/content/ko/tech/post-a.md" "$TODAY" <<'EOF'
import sys
path, today = sys.argv[1], sys.argv[2]
src = open(path).read()
src = src.replace("lastmod: 2026-04-01", "lastmod: " + today)
src = src.replace(
    "This stale body sentence still describes the situation as of early 2025 in vendor terms.",
    "The refreshed body now reflects the mid 2026 landscape with current vendor guidance and revised context.")
src = src.replace(
    "Throughput grew 40% in 2025 per the vendor study.",
    "Throughput grew 55% in 2026 per the vendor study.")
open(path, "w").write(src)
EOF
printf '{"article":"content/ko/tech/post-a.md"}\n' > "$WORK/sub.json"

# try 1 — deliberate queue-scope FAIL: an out-of-scope edit rides along
# (submit exits 0 on FAIL too — the verdict lives in the "key -> OUTCOME" line)
echo "out-of-scope line for the queue-scope fail" >> "$AGENT/content/ko/tech/post-b.md"
QUEST submit --key ko/tech/post-a --in "$WORK/sub.json" > "$RECORD/queue-scope-fail.txt" 2>&1 \
  || { cat "$RECORD/queue-scope-fail.txt" >&2; fail "submit #1 errored (not a verdict)"; }
grep -q -- '-> FAIL' "$RECORD/queue-scope-fail.txt" || fail "submit #1 did not FAIL"
grep -q 'queue-scope' "$RECORD/queue-scope-fail.txt" || fail "submit #1 did not fail on queue-scope"
git -C "$AGENT" checkout -q -- content/ko/tech/post-b.md

# try 2 — PASS
QUEST submit --key ko/tech/post-a --in "$WORK/sub.json" > "$WORK/submit-pass.txt" 2>&1 \
  || { cat "$WORK/submit-pass.txt" >&2; fail "submit #2 errored"; }
grep -q -- '-> PASS' "$WORK/submit-pass.txt" || { cat "$WORK/submit-pass.txt" >&2; fail "submit #2 did not PASS"; }
cat "$WORK/submit-pass.txt" >> "$RECORD/steps.log"
cp "$SESSION" "$RECORD/refresh-session.json"          # evidence ①
python3 - "$RECORD/refresh-session.json" <<'EOF'
import json, sys
s = json.load(open(sys.argv[1]))
it = s["items"][0]
assert it["state"] == "PASS" and it["tries"] >= 1, it
outcomes = [l["outcome"] for l in it["log"]]
assert outcomes == ["FAIL", "PASS"], outcomes
EOF

# ─────────────────────────────────────────────────────────────────────────
step "5. ① article commit → deploy pull → re-sync recognizes the new lastmod"

git -C "$AGENT" add content/ko/tech/post-a.md
git -C "$AGENT" -c user.name=agent -c user.email=agent@rehearsal.test \
  commit -q -m "refresh: ko/tech/post-a — figures advanced to 2026 (gate PASS)"
git -C "$AGENT" push -q origin main

git -C "$WORK/deploy" pull -q origin main
SYNCED=$(api POST /sites/default/sync | jsonget '["synced"]')
[ "$SYNCED" = "2" ] || fail "re-sync: synced=$SYNCED, want 2"
LASTMOD=$(api GET /sites/default/posts | python3 -c 'import json,sys
posts = json.load(sys.stdin)["posts"]
print(next(p["lastmod"] for p in posts if p["slug"] == "post-a"))')
[ "$LASTMOD" = "$TODAY" ] || fail "re-sync: post-a lastmod=$LASTMOD, want $TODAY"

# ─────────────────────────────────────────────────────────────────────────
step "6. /hooks/deployed → /archive/process → receipts done (stub)"

PLANNED=$(api POST /sites/default/hooks/deployed \
  '{"deploy_id":"rehearsal-loop1","changed":["https://fixture.example.com/tech/post-a/"]}' \
  | jsonget '["planned"]')
[ "$PLANNED" = "3" ] || fail "hooks/deployed: planned=$PLANNED, want 3"

PROCESS=$(api POST /sites/default/archive/process '{"limit":100}')
[ "$(echo "$PROCESS" | jsonget '["failed"]')" = "0" ] || fail "archive/process failed: $PROCESS"
[ "$(echo "$PROCESS" | jsonget '["deferred"]')" = "0" ] || fail "archive/process deferred: $PROCESS"

RECEIPTS=$(api GET '/sites/default/receipts?deploy_id=rehearsal-loop1&status=done')
[ "$(echo "$RECEIPTS" | jsonget '["receipts"].__len__()')" = "3" ] \
  || fail "receipts: want 3 done rows, got $RECEIPTS"
echo "$RECEIPTS" | python3 -m json.tool > "$RECORD/receipts-done.json"   # evidence ②

# ─────────────────────────────────────────────────────────────────────────
step "7. ② queue-deletion commit → export consumed sync"

git -C "$AGENT" rm -q "$QFILE"
git -C "$AGENT" -c user.name=agent -c user.email=agent@rehearsal.test \
  commit -q -m "consume: refresh-ko-tech-post-a.yaml (PASS locked, article committed)"
git -C "$AGENT" push -q origin main
git -C "$AGENT" log --oneline --reverse > "$RECORD/agent-commits.txt"

EXPORT2=$(api POST /sites/default/queue/export)
[ "$(echo "$EXPORT2" | jsonget '["consumed"]')" = "1" ] || fail "consumed sync: $EXPORT2, want consumed=1"
CONSUMED_ROWS=$(api GET '/sites/default/queue?status=consumed')
[ "$(echo "$CONSUMED_ROWS" | jsonget '["items"].__len__()')" = "1" ] \
  || fail "queue: want 1 consumed row, got $CONSUMED_ROWS"
{ echo "POST /queue/export (cycle 2):"; echo "$EXPORT2" | python3 -m json.tool
  echo; echo "GET /queue?status=consumed:"; echo "$CONSUMED_ROWS" | python3 -m json.tool
} > "$RECORD/export-consumed.json"                                       # evidence ④

# ─────────────────────────────────────────────────────────────────────────
step "8. monthly report — generate + publish to the bare origin"

REPORT=$(api POST /sites/default/reports/monthly "{\"ym\":\"$YM\"}")
[ "$(echo "$REPORT" | jsonget '["published"]')" = "True" ] || fail "report not published: $REPORT"
[ "$(echo "$REPORT" | jsonget '["articles"]')" = "2" ] || fail "report articles != 2: $REPORT"

git clone -q "file://$WORK/bare" "$WORK/report-check"
[ -f "$WORK/report-check/reports/$YM.md" ] || fail "reports/$YM.md missing in the bare origin"
grep -q "# Visibility report $YM" "$WORK/report-check/reports/$YM.md" \
  || fail "published report header drifted"
cp "$WORK/report-check/reports/$YM.md" "$RECORD/report-$YM.md"
PUBLISH_HASH=$(git -C "$WORK/report-check" log -1 --format=%H -- "reports/$YM.md")
{ echo "bare-origin publication commit for reports/$YM.md:"
  echo "$PUBLISH_HASH"
  echo
  git -C "$WORK/report-check" log -1 --stat --format='%H %ad %s' --date=iso "$PUBLISH_HASH"
  echo
  echo "full bare-origin history (agent + exporter + report bot):"
  git -C "$WORK/report-check" log --oneline --reverse
} > "$RECORD/report-publish.txt"                                          # evidence ③

# ─────────────────────────────────────────────────────────────────────────
step "9. evidence inventory"
for f in refresh-session.json receipts-done.json report-publish.txt export-consumed.json \
         queue-file.yaml next-prompt.txt queue-scope-fail.txt agent-commits.txt "report-$YM.md"; do
  [ -s "$RECORD/$f" ] || fail "evidence $f missing or empty"
  echo "  evidence: $f"
done

echo "rehearsal: PASS (record: $RECORD)"
