# Monthly visibility report scenario (Phase014). Assumes the shared-fixture
# abloqd with posts already synced (scenario-sync ran earlier in the batch
# order) and the harness psql seed applied (backend/scripts/hurl-test/run.sh
# — fixed 2026-04/2026-03 crawl_hits and gsc_snapshots rows; POST /ingest/*
# cannot place rows in a fixed past month because their cursors are
# now-relative). The explicit ym=2026-04 window is clock-free; the citation
# and queue layers of that window are deterministically empty (run_at /
# created_at are DB NOW(), which can never fall inside a fixed past month) —
# the zero answer is pinned here, non-zero goldens live in the pkg tests.
# The git publication (reports/2026-04.md in the bare origin) is asserted by
# the run.sh leg after this file passes.

# Login as the operator
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "operator@abloq.test",
  "password": "abloq-operator-test"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Generate: explicit ym — both fixture posts become report rows. The first
# generation pushes the publication copy (published=true).
POST {{host}}/reports/monthly
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": "2026-04"
}
HTTP 200
[Captures]
report_ym: jsonpath "$.ym"
[Asserts]
jsonpath "$.ym" == "2026-04"
jsonpath "$.articles" == 2
jsonpath "$.published" == true

# Read back: the stored row matches what generation reported. Seeded crawl
# layer (April window): post-a GPTBot 7+md2 + ChatGPT-User 3 = training 7,
# fetch 3; post-b ClaudeBot 4+md1. Seeded GSC: post-a 120/8, post-b 40/2.
# Weighted priority (defaults 3/1/1/2): post-a 3*3+1*7+1*120 = 136,
# post-b 1*4+1*40 = 44. March seeds (GPTBot 5, GSC 60/1) fill the trend
# column. Citation and queue layers read zero by construction.
GET {{host}}/reports/monthly/{{report_ym}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.ym" == "2026-04"
jsonpath "$.markdown" contains "# Visibility report 2026-04"
jsonpath "$.markdown" contains "| ko/tech/post-a | 2026-06-01 | 7 | 0 | 3 | 2 | 120 | 8 | 0 | 136 |"
jsonpath "$.markdown" contains "| ko/tech/post-b | 2026-06-03 | 4 | 0 | 0 | 1 | 40 | 2 | 0 | 44 |"
jsonpath "$.markdown" contains "| crawl hits | 14 | 5 |"
jsonpath "$.markdown" contains "| gsc impressions | 160 | 60 |"
jsonpath "$.markdown" contains "| gsc clicks | 10 | 1 |"
jsonpath "$.markdown" contains "| cited samples | 0 | 0 |"
jsonpath "$.markdown" contains "no queue intake in this window"
jsonpath "$.report_json" contains "\"priority\":136"
jsonpath "$.report_json" contains "\"priority\":44"
jsonpath "$.report_json" contains "\"queue\":[]"

# Regenerate the same ym: the upsert converges and the identical markdown
# makes the git publication a no-op (published=false) — idempotence.
POST {{host}}/reports/monthly
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": "2026-04"
}
HTTP 200
[Asserts]
jsonpath "$.ym" == "2026-04"
jsonpath "$.articles" == 2
jsonpath "$.published" == false

# ... and the stored row still matches (regeneration is byte-identical)
GET {{host}}/reports/monthly/{{report_ym}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.markdown" contains "| ko/tech/post-a | 2026-06-01 | 7 | 0 | 3 | 2 | 120 | 8 | 0 | 136 |"

# A month nobody generated is a plain 404 (guard path)
GET {{host}}/reports/monthly/{{missing_ym}}
Authorization: Bearer {{token}}
[Options]
variable: missing_ym="1999-01"
HTTP 404
