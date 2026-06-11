#!/usr/bin/env bash
# Full Hurl harness on a throwaway postgres cluster (Phase012·020).
#
# The system postgres (5432) is not usable in the dev sandbox, so this
# script initdb's a temporary cluster on 127.0.0.1:55432 (trust auth),
# applies the generated migrations, seeds the test operator/viewer, starts
# the archive stub plus abloqd (the listen address is fixed at :8080, so the
# multisite instance and the cluster-fixture instance run one after the
# other), runs every Hurl file in the shared-DB order (scenario-freshness
# first: its cold-start priority assert needs empty crawl_hits AND empty
# queue_items; scenario-crawl after it for the same reason;
# scenario-gsc-citation before smoke — smoke re-runs its ops loosely;
# scenario-report after a fixed-month psql seed — POST /ingest/* cursors are
# now-relative and cannot land rows in a fixed past month, so the 2026-04
# crawl/GSC rows are seeded directly (users seed precedent);
# scenario-multisite after scenario-report — it needs default's posts,
# queue rows and the seeded 2026-04 report while site beta stays untouched;
# smoke last), verifies the report publication copy in the bare origin, and
# tears everything down. Login is rate-limited 5/min/IP, so the run sleeps
# 61s between the login batches (4+3+3 logins). --jobs 1 pins the file
# order (hurl --test parallelizes by default since 4.2).
#
# Phase020 multisite: the shared instance boots from a 2-site+1-inactive
# sites.yaml fixture (SITES_YAML_PATH) — site `default` is the shared
# fixture blog, site `beta` a copy of it (same slugs on purpose: the
# coexistence oracle), site `retired` is declared active: false. The
# cluster instance keeps the BLOG_REPO_PATH-only boot, pinning the
# single-site default-synthesis backward-compat path.
#
# Usage: backend/scripts/hurl-test/run.sh
# Requires: backend/arts generated (yongol generate backend/specs backend/arts)
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../../.." && pwd)"
PG="${PG_BIN:-/usr/lib/postgresql/16/bin}"
WORK="$(mktemp -d)"
PGPORT=55432
PORT_STUB=18099
HOST="http://127.0.0.1:8080"
PIDS=()
ABLOQD_PID=""

cleanup() {
  [ -n "$ABLOQD_PID" ] && kill "$ABLOQD_PID" 2>/dev/null || true
  for pid in "${PIDS[@]:-}"; do kill "$pid" 2>/dev/null || true; done
  "$PG/pg_ctl" -D "$WORK/pg" stop -m immediate >/dev/null 2>&1 || true
  rm -rf "$WORK"
}
trap cleanup EXIT

# ① temporary postgres cluster, trust auth, 127.0.0.1:55432
"$PG/initdb" -D "$WORK/pg" -A trust -U postgres >/dev/null
"$PG/pg_ctl" -D "$WORK/pg" -o "-p $PGPORT -k $WORK/pg -c listen_addresses=127.0.0.1" \
  -l "$WORK/pg.log" start >/dev/null
for db in abloqd abloqd_cluster; do
  "$PG/createdb" -h 127.0.0.1 -p "$PGPORT" -U postgres "$db"
  for mig in "$ROOT"/backend/arts/db/migrations/*.up.sql; do
    "$PG/psql" -q -h 127.0.0.1 -p "$PGPORT" -U postgres -d "$db" -f "$mig" >/dev/null
  done
done

# ② seed the Hurl users (bcrypt via the arts module — x/crypto is a dep)
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
OP_HASH=$(cd "$ROOT/backend/arts/backend" && go run "$WORK/hashgen.go" abloq-operator-test)
VW_HASH=$(cd "$ROOT/backend/arts/backend" && go run "$WORK/hashgen.go" abloq-viewer-test)
for db in abloqd abloqd_cluster; do
  "$PG/psql" -q -h 127.0.0.1 -p "$PGPORT" -U postgres -d "$db" <<SQL
INSERT INTO users (email, password_hash, role)
VALUES ('operator@abloq.test', '$OP_HASH', 'operator'),
       ('viewer@abloq.test', '$VW_HASH', 'viewer');
SQL
done

# ③ server binary, archive stub, bare origins, throwaway GSC SA JSON
(cd "$ROOT/backend/arts/backend" && go build -o "$WORK/abloqd" ./cmd)
python3 "$ROOT/backend/scripts/archive-stub/archive_stub.py" "$PORT_STUB" &
PIDS+=($!)
"$ROOT/backend/scripts/queue-export-test/setup-bare.sh" "$WORK/bare-a" >/dev/null
"$ROOT/backend/scripts/queue-export-test/setup-bare.sh" "$WORK/bare-beta" >/dev/null
"$ROOT/backend/scripts/queue-export-test/setup-bare.sh" "$WORK/bare-b" >/dev/null

# site beta — a copy of the shared fixture blog: same (lang, section, slug)
# keys on purpose, the multisite same-slug coexistence oracle.
cp -r "$ROOT/backend/fixtures/blog" "$WORK/blog-beta"

# the multisite SSOT fixture: 2 active sites + 1 declared-inactive site.
# Empty per-site keys (indexnow_key, gsc.sa_json_path) fall back to the
# instance env below — the fallback path is part of what this harness pins.
cat > "$WORK/sites.yaml" <<EOF
sites:
  - name: default
    repo_path: $ROOT/backend/fixtures/blog
    queue_export:
      repo_url: file://$WORK/bare-a
    cf_log_source: $ROOT/backend/fixtures/cflogs
  - name: beta
    repo_path: $WORK/blog-beta
    queue_export:
      repo_url: file://$WORK/bare-beta
  - name: retired
    repo_path: $ROOT/backend/fixtures/blog
    active: false
EOF

openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out "$WORK/sa.pem" 2>/dev/null
python3 - "$WORK/sa.pem" "$WORK/sa.json" <<'EOF'
import json, sys
json.dump({"client_email": "stub@test-project.iam.gserviceaccount.com",
           "private_key": open(sys.argv[1]).read()}, open(sys.argv[2], "w"))
EOF

start_abloqd() { # $1=db $2=blog-repo|sites:<path> $3=bare $4=export-workdir-base $5=logfile
  # $2 selects the boot mode: "sites:<path>" boots from the sites.yaml SSOT
  # (multisite), a plain path boots the BLOG_REPO_PATH default synthesis
  # (single-site backward compat — the cluster instance pins it).
  local repo_env=() 
  case "$2" in
    sites:*) repo_env=(SITES_YAML_PATH="${2#sites:}");;
    *) repo_env=(BLOG_REPO_PATH="$2" QUEUE_EXPORT_REPO_URL="file://$3");;
  esac
  env DATABASE_URL="postgres://postgres@127.0.0.1:$PGPORT/$1?sslmode=disable" \
    JWT_SECRET="hurl-test-secret-hurl-test-secret-0123456789" \
    OPA_POLICY_PATH="$ROOT/backend/arts/backend/policy/authz.rego" \
    GIN_MODE=release LOG_LEVEL=error \
    WAYBACK_BASE_URL="http://127.0.0.1:$PORT_STUB" \
    INDEXNOW_ENDPOINT="http://127.0.0.1:$PORT_STUB/indexnow" \
    INDEXNOW_KEY="hurltestkey" \
    GSC_API_BASE="http://127.0.0.1:$PORT_STUB" \
    GSC_TOKEN_URL="http://127.0.0.1:$PORT_STUB/token" \
    GSC_SA_JSON_PATH="$WORK/sa.json" \
    LINKCHECK_HOST_OVERRIDE="http://127.0.0.1:$PORT_STUB" \
    GSC_SEARCH_API_BASE="http://127.0.0.1:$PORT_STUB" \
    GSC_LOOKBACK_DAYS=3 \
    GSC_INSPECT_RECENT_DAYS=36500 \
    PERPLEXITY_API_KEY="stub-key" PERPLEXITY_BASE_URL="http://127.0.0.1:$PORT_STUB" \
    OPENAI_API_KEY="stub-key" OPENAI_BASE_URL="http://127.0.0.1:$PORT_STUB" \
    ANTHROPIC_API_KEY="stub-key" ANTHROPIC_BASE_URL="http://127.0.0.1:$PORT_STUB" \
    CITATION_INTERVAL_MS=0 \
    "${repo_env[@]}" \
    QUEUE_EXPORT_WORKDIR="$4" \
    CF_LOG_SOURCE="$ROOT/backend/fixtures/cflogs" \
    "$WORK/abloqd" > "$WORK/$5" 2>&1 &
  ABLOQD_PID=$!
  for _ in $(seq 1 100); do
    curl -sf "$HOST/health" >/dev/null 2>&1 && return 0
    sleep 0.2
  done
  echo "abloqd did not come up — $WORK/$5:" >&2
  tail -20 "$WORK/$5" >&2
  return 1
}

stop_abloqd() {
  kill "$ABLOQD_PID" 2>/dev/null || true
  wait "$ABLOQD_PID" 2>/dev/null || true
  ABLOQD_PID=""
}

TESTS="$ROOT/backend/arts/tests"

# ④ shared multisite instance (sites.yaml SSOT): freshness first
# (cold-start oracle), multisite after report, smoke last.
# Login rate limit is 5/min/IP — batches of <=4 logins, 61s apart.
start_abloqd abloqd "sites:$WORK/sites.yaml" "$WORK/bare-a" "$WORK/export" abloqd-a.log
hurl --test --jobs 1 --variable "host=$HOST" \
  "$TESTS/scenario-freshness.hurl" \
  "$TESTS/scenario-evidence.hurl" \
  "$TESTS/scenario-sync.hurl" \
  "$TESTS/scenario-archive.hurl"
echo "sleeping 61s (login rate limit window)..."
sleep 61
hurl --test --jobs 1 --variable "host=$HOST" \
  "$TESTS/scenario-auth-forbidden.hurl" \
  "$TESTS/scenario-crawl.hurl" \
  "$TESTS/scenario-gsc-citation.hurl"
echo "sleeping 61s (login rate limit window)..."
sleep 61

# fixed-month seed for scenario-report (2026-04 window + 2026-03 trend):
# deterministic forever — NOW() can never fall inside a fixed past month,
# so the citation/queue layers of that window stay zero by construction.
# rows belong to site `default` — the id rides in by name (never hardcoded)
"$PG/psql" -q -h 127.0.0.1 -p "$PGPORT" -U postgres -d abloqd <<'SQL'
INSERT INTO crawl_hits (site_id, hit_date, bot, lang, section, slug, hits, md_hits) VALUES
  ((SELECT id FROM sites WHERE name='default'), '2026-04-10', 'GPTBot',       'ko', 'tech', 'post-a', 7, 2),
  ((SELECT id FROM sites WHERE name='default'), '2026-04-12', 'ChatGPT-User', 'ko', 'tech', 'post-a', 3, 0),
  ((SELECT id FROM sites WHERE name='default'), '2026-04-15', 'ClaudeBot',    'ko', 'tech', 'post-b', 4, 1),
  ((SELECT id FROM sites WHERE name='default'), '2026-03-20', 'GPTBot',       'ko', 'tech', 'post-a', 5, 0);
INSERT INTO gsc_snapshots (site_id, snap_date, page, impressions, clicks, avg_position) VALUES
  ((SELECT id FROM sites WHERE name='default'), '2026-04-10', 'https://fixture.example.com/tech/post-a/', 120, 8, 3.2),
  ((SELECT id FROM sites WHERE name='default'), '2026-04-11', 'https://fixture.example.com/tech/post-b/',  40, 2, 7.5),
  ((SELECT id FROM sites WHERE name='default'), '2026-03-15', 'https://fixture.example.com/tech/post-a/',  60, 1, 9.9);
SQL

hurl --test --jobs 1 --variable "host=$HOST" \
  "$TESTS/scenario-report.hurl" \
  "$TESTS/scenario-multisite.hurl" \
  "$TESTS/smoke.hurl"
stop_abloqd

# git publication leg: the report markdown landed in the bare origin as a
# publication copy (the DB row stays the lookup truth).
git clone -q "file://$WORK/bare-a" "$WORK/report-check"
grep -q '# Visibility report 2026-04' "$WORK/report-check/reports/2026-04.md" \
  || { echo "FAIL: reports/2026-04.md missing or wrong in the bare origin" >&2; exit 1; }
grep -q '| ko/tech/post-a | 2026-06-01 | 7 | 0 | 3 | 2 | 120 | 8 | 0 | 136 |' \
  "$WORK/report-check/reports/2026-04.md" \
  || { echo "FAIL: published report row drifted" >&2; exit 1; }
echo "ok: report publication copy verified in the bare origin"

# site beta's report publication landed in ITS OWN bare origin — the
# per-site work-clone/origin isolation leg of the multisite scenario.
git clone -q "file://$WORK/bare-beta" "$WORK/report-check-beta"
grep -q '# Visibility report 2026-04' "$WORK/report-check-beta/reports/2026-04.md" \
  || { echo "FAIL: site beta reports/2026-04.md missing in its bare origin" >&2; exit 1; }
echo "ok: site beta report publication isolated in its own origin"

# ⑤ dedicated cluster instance (own DB, own fixture, fresh limiter) —
# cluster oracle + the citation budget=0 no-op oracle (citation_budget
# defaults to 0 in the cluster-blog fixture). Boots WITHOUT SITES_YAML_PATH:
# BLOG_REPO_PATH alone must synthesize the `default` site (backward compat).
start_abloqd abloqd_cluster "$ROOT/backend/fixtures/cluster-blog" "$WORK/bare-b" "$WORK/export-b" abloqd-b.log
hurl --test --jobs 1 --variable "host=$HOST" \
  "$TESTS/scenario-cluster.hurl" \
  "$TESTS/scenario-citation-budget0.hurl"
stop_abloqd

echo "hurl-test: PASS"
