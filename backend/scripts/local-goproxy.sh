#!/usr/bin/env bash
# Build a file:// GOPROXY that serves the abloq working tree as
# github.com/park-jun-woo/abloq@v0.0.1.
#
# Why: backend/specs/func/content/index_repo.go imports
# github.com/park-jun-woo/abloq/pkg/content, and `yongol generate` runs
# `go mod tidy` inside the generated backend. Until pkg/content is pushed
# upstream the module proxy cannot resolve it, so we publish the working
# tree into a local file proxy and put it in front of proxy.golang.org.
#
# Usage:
#   backend/scripts/local-goproxy.sh            # builds /tmp/abloq-goproxy
#   export GOPROXY="file:///tmp/abloq-goproxy,https://proxy.golang.org,direct"
#   export GONOSUMDB="github.com/park-jun-woo/abloq"
#   yongol generate backend/specs backend/arts
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
VERSION="${1:-v0.0.1}"
PROXY="${ABLOQ_GOPROXY_DIR:-/tmp/abloq-goproxy}"
MOD="github.com/park-jun-woo/abloq"
VDIR="$PROXY/$MOD/@v"
STAGE="$(mktemp -d)"
trap 'rm -rf "$STAGE"' EXIT

mkdir -p "$VDIR" "$STAGE/$MOD@$VERSION"
cp "$ROOT/go.mod" "$ROOT/go.sum" "$STAGE/$MOD@$VERSION/"
cp -r "$ROOT/pkg" "$STAGE/$MOD@$VERSION/pkg"

python3 - "$STAGE" "$VDIR/$VERSION.zip" <<'EOF'
import os, sys, zipfile
stage, out = sys.argv[1], sys.argv[2]
with zipfile.ZipFile(out, "w", zipfile.ZIP_DEFLATED) as z:
    for base, _, files in os.walk(stage):
        for f in sorted(files):
            p = os.path.join(base, f)
            z.write(p, os.path.relpath(p, stage))
EOF

cp "$ROOT/go.mod" "$VDIR/$VERSION.mod"
printf '{"Version":"%s","Time":"2026-01-01T00:00:00Z"}\n' "$VERSION" > "$VDIR/$VERSION.info"
printf '%s\n' "$VERSION" > "$VDIR/list"

echo "local GOPROXY ready: $PROXY"
echo "  export GOPROXY=\"file://$PROXY,https://proxy.golang.org,direct\""
echo "  export GONOSUMDB=\"$MOD\""
