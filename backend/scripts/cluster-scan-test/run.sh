#!/usr/bin/env bash
# Candidate-existence + CLI equivalence legs for the cluster scenario
# (Phase011).
#
# queue_items.payload is @sensitive (the QueueItem response never carries
# it), so Hurl cannot assert the candidates payload. This script runs AFTER
# scenario-cluster.hurl passed (4 rows exported, files pushed to the bare
# origin):
#   1. clone the bare origin, check all 4 cluster queue files exist and each
#      carries the gate-contract `key:` line plus the violations and
#      candidates payload keys,
#   2. run the abloq CLI on a scratch copy of the dedicated fixture and
#      `diff -r` its quests/queue/ against the exported clone's — the
#      serialization is deterministic, so the trees must be byte-identical.
#
# Usage: run.sh <bare-path> <fixture-path>
#   bare-path     the setup-bare.sh path the second abloqd's
#                 QUEUE_EXPORT_REPO_URL points at
#   fixture-path  backend/fixtures/cluster-blog (copied, never written)
#
# Env: ABLOQ_BIN — abloq binary (default: abloq on PATH).
set -euo pipefail

BARE="${1:?usage: run.sh <bare-path> <fixture-path>}"
FIXTURE="${2:?usage: run.sh <bare-path> <fixture-path>}"
ABLOQ_BIN="${ABLOQ_BIN:-abloq}"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

fail() { echo "FAIL: $*" >&2; exit 1; }

# ① Every exported cluster queue file carries key + violations + candidates
git clone -q "file://$BARE" "$WORK/clone"
QUEUE_DIR="$WORK/clone/quests/queue"
COUNT=$(find "$QUEUE_DIR" -name 'cluster-*.yaml' | wc -l)
[ "$COUNT" -eq 4 ] || fail "expected 4 cluster queue files, found $COUNT"
for f in "$QUEUE_DIR"/cluster-*.yaml; do
  grep -q '^key: ' "$f" || fail "$f has no gate-contract key: line"
  grep -q '^  violations: ' "$f" || fail "$f has no violations payload"
  grep -q '^  candidates: ' "$f" || fail "$f has no candidates payload"
done
echo "ok: 4 cluster queue files, each carries key/violations/candidates"

# ② CLI equivalence: stateless local scan produces byte-identical files
cp -r "$FIXTURE" "$WORK/fixture"
"$ABLOQ_BIN" scan cluster "$WORK/fixture" >/dev/null
diff -r "$WORK/fixture/quests/queue" "$QUEUE_DIR" \
  || fail "CLI queue files differ from the exported clone"
echo "ok: CLI quests/queue and exported clone are byte-identical (diff -r 0)"

echo "cluster-scan-test: PASS"
