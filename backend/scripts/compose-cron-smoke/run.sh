#!/usr/bin/env bash
# Shell-expansion smoke for every compose cron service (Phase019).
#
# Why this exists: `docker compose config` validates YAML, not shell. The
# cron commands are `sh -ec` scripts whose JSON bodies depend on correct
# quote escaping ( \" inside a double-quoted -d argument ). Phase019 review
# found a live bug: gsc-ingest and citation-sample had lost the backslashes,
# so their login body expanded to {email:...} — invalid JSON, silent login
# failure every cycle. This smoke would have caught it, and now pins it.
#
# How: resolve the compose file with every profile active, extract each
# profiled service's command script and environment, then run the script in
# a sandbox PATH where apk/curl/jq/sleep are stubs:
#   - curl records every call (url/method/headers/-d body) and answers
#     /auth/login with a fixed token JSON,
#   - jq extracts access_token like the real thing,
#   - sleep returns 0 once then 86, so `while true` runs exactly one
#     iteration and the `sh -ec` exits 86 (the expected terminator).
# Then assert per service: the login body parses as JSON and carries the
# operator credentials verbatim, every follow-up call hits the expected
# endpoint with the Bearer token, and every JSON body parses.
#
# Usage: backend/scripts/compose-cron-smoke/run.sh
# Requires: docker compose v2 (client-side only — no daemon needed), python3
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../../.." && pwd)"
COMPOSE="$ROOT/deploy/backend/docker-compose.yaml"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

# ① resolve the compose file with all profiles + dummy required env.
# This is also the `docker compose config` gate: a YAML/interpolation error
# fails here.
export POSTGRES_PASSWORD=smoke-pg-password
export JWT_SECRET=smoke-jwt-secret-smoke-jwt-secret-0123
export BLOG_REPO_PATH=/tmp
export ABLOQD_OPERATOR_EMAIL=operator@smoke.test
# NB: the cron scripts splice the password into a JSON body verbatim, so an
# operator password must not contain `"` or `\` (documented in .env.example
# and docs/operations.md — same constraint as the CI hook in
# template/files/deploy/archiver.md). Specials other than those are fine:
export ABLOQD_OPERATOR_PASSWORD='smoke-p@ss!word $123'
export COMPOSE_PROFILES=backstop,crawl,gsc,citation,report,freshness,evidence,cluster,queue
docker compose -f "$COMPOSE" config --format json > "$WORK/config.json"
echo "ok: docker compose config resolved ($(wc -c < "$WORK/config.json") bytes)"

# ② stub PATH
STUB="$WORK/stub"
mkdir -p "$STUB"

cat > "$STUB/apk" <<'EOF'
#!/bin/sh
exit 0
EOF

cat > "$STUB/curl" <<'EOF'
#!/bin/sh
# record one call: url / method / headers / -d body, then answer.
n=$(cat "$SMOKE_DIR/count" 2>/dev/null || echo 0)
n=$((n+1)); echo "$n" > "$SMOKE_DIR/count"
out="$SMOKE_DIR/call-$n"
url=""
while [ $# -gt 0 ]; do
  case "$1" in
    -d|--data) shift; printf '%s' "$1" > "$out.data";;
    -H|--header) shift; printf '%s\n' "$1" >> "$out.headers";;
    -X) shift; printf '%s' "$1" > "$out.method";;
    -m|--max-time) shift;;
    http*) url="$1"; printf '%s' "$1" > "$out.url";;
  esac
  shift
done
case "$url" in
  */auth/login) printf '%s' '{"access_token":"smoke-token"}';;
  */sites?active_filter=true) printf '%s' '{"sites":[{"name":"default"}]}';;
esac
exit 0
EOF

cat > "$STUB/jq" <<'EOF'
#!/bin/sh
# emulate `jq -r .access_token` and `jq -r '.sites[].name'`
case "$2" in
  .access_token) sed -n 's/.*"access_token"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p';;
  .sites*) sed -n 's/.*"name"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p';;
  *) cat;;
esac
EOF

cat > "$STUB/sleep" <<'EOF'
#!/bin/sh
# one loop iteration: first call proceeds, second call terminates sh -ec
if [ -f "$SMOKE_DIR/slept" ]; then exit 86; fi
: > "$SMOKE_DIR/slept"
exit 0
EOF

chmod +x "$STUB"/*

# ③ run every profiled (cron) service command under the stubs and assert
python3 - "$WORK/config.json" "$STUB" "$WORK" <<'EOF'
import json, os, subprocess, sys

cfg_path, stub, work = sys.argv[1], sys.argv[2], sys.argv[3]
cfg = json.load(open(cfg_path))

# expected follow-up endpoints (in call order, after the login call).
# Phase020: every cron fans out over GET /sites?active_filter=true and calls
# the per-site op for each name — the stub registry answers one site
# (`default`), so the expected calls are the list lookup plus one site round.
EXPECT = {
    "archiver-backstop": ["/sites?active_filter=true", "/sites/default/receipts/retry", "/sites/default/archive/process"],
    "crawl-ingest":      ["/sites?active_filter=true", "/sites/default/ingest/crawl"],
    "gsc-ingest":        ["/sites?active_filter=true", "/sites/default/ingest/gsc"],
    "citation-sample":   ["/sites?active_filter=true", "/sites/default/sample/citations"],
    "report-monthly":    ["/sites?active_filter=true", "/sites/default/reports/monthly"],
    "freshness-scan":    ["/sites?active_filter=true", "/sites/default/scans/freshness"],
    "evidence-scan":     ["/sites?active_filter=true", "/sites/default/scans/evidence"],
    "cluster-scan":      ["/sites?active_filter=true", "/sites/default/scans/cluster"],
    "queue-export":      ["/sites?active_filter=true", "/sites/default/queue/export"],
}

services = cfg["services"]
cron = {n: s for n, s in services.items() if s.get("profiles")}
missing = sorted(set(EXPECT) - set(cron))
extra = sorted(set(cron) - set(EXPECT))
if missing or extra:
    sys.exit(f"FAIL: cron service set drifted — missing={missing} extra={extra}")

failures = []
for name, svc in sorted(cron.items()):
    cmd = svc.get("command")
    if not (isinstance(cmd, list) and len(cmd) == 3 and cmd[0] == "sh" and cmd[1] == "-ec"):
        failures.append(f"{name}: command is not [sh, -ec, script]: {cmd!r}")
        continue
    # `config` re-escapes $ as $$ so its output stays a valid compose file;
    # the container runtime receives a single $ — replicate that here.
    script = cmd[2].replace("$$", "$")
    smoke_dir = os.path.join(work, "run-" + name)
    os.makedirs(smoke_dir)
    env = dict(os.environ)
    env.update({k: str(v) for k, v in (svc.get("environment") or {}).items()})
    env["PATH"] = stub + ":" + env["PATH"]
    env["SMOKE_DIR"] = smoke_dir
    p = subprocess.run(["sh", "-ec", script], env=env,
                       capture_output=True, text=True, timeout=30)
    if p.returncode != 86:
        failures.append(f"{name}: exit {p.returncode} (want 86 = one full iteration)"
                        f" stderr={p.stderr.strip()!r} stdout={p.stdout.strip()!r}")
        continue

    def call(i, part):
        path = os.path.join(smoke_dir, f"call-{i}.{part}")
        return open(path).read() if os.path.exists(path) else ""

    n = int(call_count := open(os.path.join(smoke_dir, "count")).read().strip())
    want = 1 + len(EXPECT[name])
    if n != want:
        failures.append(f"{name}: {n} curl calls (want {want})")
        continue

    # call 1: login — the escape regression oracle
    if not call(1, "url").endswith("/auth/login"):
        failures.append(f"{name}: first call is {call(1,'url')!r}, want /auth/login")
        continue
    body = call(1, "data")
    try:
        login = json.loads(body)
    except ValueError:
        failures.append(f"{name}: login body is not JSON (escape bug): {body!r}")
        continue
    if login.get("email") != env["ABLOQD_OPERATOR_EMAIL"] or \
       login.get("password") != env["ABLOQD_OPERATOR_PASSWORD"]:
        failures.append(f"{name}: login credentials drifted: {login!r}")
        continue

    # follow-up calls: endpoint order, bearer token, JSON bodies
    for i, suffix in enumerate(EXPECT[name], start=2):
        url, headers, data = call(i, "url"), call(i, "headers"), call(i, "data")
        if not url.endswith(suffix):
            failures.append(f"{name}: call {i} is {url!r}, want *{suffix}")
            break
        if "Authorization: Bearer smoke-token" not in headers:
            failures.append(f"{name}: call {i} lacks the bearer token: {headers!r}")
            break
        if data:
            try:
                json.loads(data)
            except ValueError:
                failures.append(f"{name}: call {i} body is not JSON: {data!r}")
                break
    else:
        print(f"ok: {name} — login JSON valid, " +
              " -> ".join(EXPECT[name]) + " with bearer token")

if failures:
    print("\n".join("FAIL: " + f for f in failures), file=sys.stderr)
    sys.exit(1)
EOF

echo "compose-cron-smoke: PASS"
