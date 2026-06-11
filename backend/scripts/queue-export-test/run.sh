#!/usr/bin/env bash
# Consumed-sync driver for the freshness scenario (Phase009).
#
# Hurl is HTTP-only, so the git legs of the scenario live here. Run AFTER
# scenario-freshness.hurl passed (2 rows exported, files pushed to the bare
# origin):
#   1. clone the bare origin, check both queue files carry the gate-contract
#      `key:` line,
#   2. consume one item the way an agent does — delete its queue file,
#      commit, push,
#   3. POST /queue/export and assert the row synced exported → consumed.
#
# Usage: run.sh <host> <bare-path>
#   host       e.g. http://127.0.0.1:18080
#   bare-path  the setup-bare.sh path abloqd's QUEUE_EXPORT_REPO_URL points at
#
# Credentials: ABLOQD_OPERATOR_EMAIL / ABLOQD_OPERATOR_PASSWORD env
# (defaults: the Hurl test operator).
set -euo pipefail

HOST="${1:?usage: run.sh <host> <bare-path>}"
BARE="${2:?usage: run.sh <host> <bare-path>}"
EMAIL="${ABLOQD_OPERATOR_EMAIL:-operator@abloq.test}"
PASSWORD="${ABLOQD_OPERATOR_PASSWORD:-abloq-operator-test}"
SITE="${ABLOQD_SITE:-default}"   # Phase020: domain ops live under /sites/<name>/
CONSUME_FILE="${CONSUME_FILE:-refresh-ko-tech-post-b.yaml}"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

fail() { echo "FAIL: $*" >&2; exit 1; }

# ① The exported queue files satisfy the gate contract (key: <lang>/<section>/<slug>)
git clone -q "file://$BARE" "$WORK/clone"
QUEUE_DIR="$WORK/clone/quests/queue"
COUNT=$(find "$QUEUE_DIR" -name '*.yaml' | wc -l)
[ "$COUNT" -ge 2 ] || fail "expected >= 2 queue files, found $COUNT"
for f in "$QUEUE_DIR"/*.yaml; do
  grep -q '^key: ' "$f" || fail "$f has no gate-contract key: line"
done
echo "ok: $COUNT queue files, every file carries a key: line"

# ② Agent consumption commit: delete one queue file and push.
#    (Real agents update the article first, then delete in a separate commit —
#    template/files docs; here only the deletion matters to the sync.)
[ -f "$QUEUE_DIR/$CONSUME_FILE" ] || fail "$CONSUME_FILE not in the queue"
git -C "$WORK/clone" rm -q "quests/queue/$CONSUME_FILE"
git -C "$WORK/clone" -c user.name=agent -c user.email=agent@test \
  commit -q -m "consume $CONSUME_FILE"
git -C "$WORK/clone" push -q origin HEAD
echo "ok: agent deleted $CONSUME_FILE and pushed"

# ③ The next export cycle syncs the deletion: exported → consumed
jsonget() { python3 -c 'import json,sys; print(json.load(sys.stdin)'"$1"')'; }

TOKEN=$(curl -sf -X POST "$HOST/auth/login" -H 'Content-Type: application/json' \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}" | jsonget '["access_token"]')
[ -n "$TOKEN" ] || fail "login failed"

EXPORT=$(curl -sf -X POST "$HOST/sites/$SITE/queue/export" -H "Authorization: Bearer $TOKEN")
CONSUMED=$(echo "$EXPORT" | jsonget '["consumed"]')
[ "$CONSUMED" = "1" ] || fail "expected consumed=1, got: $EXPORT"
echo "ok: export reported consumed=1"

ROWS=$(curl -sf "$HOST/sites/$SITE/queue?status=consumed" -H "Authorization: Bearer $TOKEN" \
  | jsonget '["items"].__len__()')
[ "$ROWS" = "1" ] || fail "expected 1 consumed row, got $ROWS"
echo "ok: 1 row is status=consumed"

echo "queue-export-test: PASS"
