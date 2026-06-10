#!/usr/bin/env bash
# Create the local bare origin the queue exporter pushes to in tests.
#
# Builds <bare-path> as a bare git repository whose main branch holds one
# seed commit (a clone must have a HEAD for pull --rebase). Point abloqd at
# it with:
#   QUEUE_EXPORT_REPO_URL=file://<bare-path>
#   QUEUE_EXPORT_WORKDIR=<scratch dir>
#
# Usage: setup-bare.sh <bare-path>
set -euo pipefail

BARE="${1:?usage: setup-bare.sh <bare-path>}"
SEED="$(mktemp -d)"
trap 'rm -rf "$SEED"' EXIT

rm -rf "$BARE"
git init --bare -b main "$BARE" >/dev/null
git init -b main "$SEED" >/dev/null
echo "queue export test fixture" > "$SEED/README.md"
git -C "$SEED" add README.md
git -C "$SEED" -c user.name=seed -c user.email=seed@test commit -q -m seed
git -C "$SEED" push -q "file://$BARE" main

echo "bare origin ready: file://$BARE"
