#!/usr/bin/env bash
# Build a file:// GOPROXY that serves three local modules:
#   github.com/park-jun-woo/abloq   @ $1 (default v0.0.11) — staged from the
#     working tree (go.mod/go.sum/pkg only; bump on every pkg/ change: the go
#     module index caches same-version republish)
#   github.com/park-jun-woo/toulmin @ v0.1.0 — staged from the git tag
#   github.com/park-jun-woo/reins   @ v0.2.0 — staged from the git tag
#
# Why: backend/specs imports github.com/park-jun-woo/abloq/pkg/content, and
# the abloq root module requires github.com/park-jun-woo/reins (Phase016
# writing quest), which requires toulmin. None of the three is published, so
# we publish them into a local file proxy in front of proxy.golang.org.
#
# reins/toulmin are staged with `git archive` at the tag (NOT cp -r): the
# working trees carry untracked junk (toulmin py/.venv) and nested modules
# (reins ccnews/, comail/) that a module zip must not contain.
#
# Usage:
#   backend/scripts/local-goproxy.sh            # builds /tmp/abloq-goproxy
#   export GOPROXY="file:///tmp/abloq-goproxy,https://proxy.golang.org,direct"
#   export GONOSUMDB="github.com/park-jun-woo/abloq,github.com/park-jun-woo/reins,github.com/park-jun-woo/toulmin"
#   export GOFLAGS=-mod=mod
#   yongol generate backend/specs backend/arts
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
VERSION="${1:-v0.0.11}"
PROXY="${ABLOQ_GOPROXY_DIR:-/tmp/abloq-goproxy}"
REINS_DIR="${REINS_DIR:-$ROOT/../reins}"
TOULMIN_DIR="${TOULMIN_DIR:-$ROOT/../toulmin}"
STAGE="$(mktemp -d)"
trap 'rm -rf "$STAGE"' EXIT

zipstage() { # $1=stage root  $2=output zip
  python3 - "$1" "$2" <<'EOF'
import os, sys, zipfile
stage, out = sys.argv[1], sys.argv[2]
with zipfile.ZipFile(out, "w", zipfile.ZIP_DEFLATED) as z:
    for base, _, files in os.walk(stage):
        for f in sorted(files):
            p = os.path.join(base, f)
            z.write(p, os.path.relpath(p, stage))
EOF
}

writemeta() { # $1=@v dir  $2=version  $3=staged go.mod
  cp "$3" "$1/$2.mod"
  printf '{"Version":"%s","Time":"2026-01-01T00:00:00Z"}\n' "$2" > "$1/$2.info"
  (cd "$1" && ls *.info | sed 's/\.info$//' > list)
}

# --- abloq: working tree (go.mod/go.sum/pkg only) -------------------------
MOD="github.com/park-jun-woo/abloq"
VDIR="$PROXY/$MOD/@v"
mkdir -p "$VDIR" "$STAGE/$MOD@$VERSION"
cp "$ROOT/go.mod" "$ROOT/go.sum" "$STAGE/$MOD@$VERSION/"
cp -r "$ROOT/pkg" "$STAGE/$MOD@$VERSION/pkg"
zipstage "$STAGE" "$VDIR/$VERSION.zip"
writemeta "$VDIR" "$VERSION" "$STAGE/$MOD@$VERSION/go.mod"
rm -rf "$STAGE/${MOD%%/*}"

# --- reins/toulmin: git archive at the tag, nested modules stripped -------
tagmodule() { # $1=repo dir  $2=module path  $3=tag
  local dir="$1" mod="$2" tag="$3" vdir="$PROXY/$2/@v"
  mkdir -p "$vdir" "$STAGE/$mod@$tag"
  git -C "$dir" archive --format=tar "$tag" | tar -x -C "$STAGE/$mod@$tag"
  # a module zip must not contain nested modules (e.g. reins ccnews/, comail/)
  find "$STAGE/$mod@$tag" -mindepth 2 -name go.mod -printf '%h\n' \
    | xargs -r rm -rf
  zipstage "$STAGE" "$vdir/$tag.zip"
  writemeta "$vdir" "$tag" "$STAGE/$mod@$tag/go.mod"
  rm -rf "$STAGE/${mod%%/*}"
}

tagmodule "$TOULMIN_DIR" "github.com/park-jun-woo/toulmin" "v0.1.0"
if git -C "$REINS_DIR" rev-parse -q --verify "refs/tags/v0.2.0" >/dev/null; then
  tagmodule "$REINS_DIR" "github.com/park-jun-woo/reins" "v0.2.0"
else
  echo "WARN: reins tag v0.2.0 missing — reins not staged (bootstrap order: tag toulmin, tidy+tag reins, re-run)"
fi

echo "local GOPROXY ready: $PROXY"
echo "  export GOPROXY=\"file://$PROXY,https://proxy.golang.org,direct\""
echo "  export GONOSUMDB=\"github.com/park-jun-woo/abloq,github.com/park-jun-woo/reins,github.com/park-jun-woo/toulmin\""
