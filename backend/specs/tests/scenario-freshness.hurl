# Freshness scanner + queue exporter scenario — scan → enqueue → re-scan
# idempotence (zero duplicates) → export to the local bare repo → idempotent
# re-export. Assumes a fresh database (no queue_items rows yet) and an abloqd
# started with QUEUE_EXPORT_REPO_URL/QUEUE_EXPORT_WORKDIR pointing at a local
# bare repo (backend/scripts/queue-export-test). The consumed-sync leg needs
# git surgery between requests, so backend/scripts/queue-export-test/run.sh
# drives it after this file passes.

# Login as the operator
POST {{host}}/auth/login
Content-Type: application/json
[Options]
variable: site="default"
{
  "email": "operator@abloq.test",
  "password": "abloq-operator-test"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Sync the fixture repository — the scanner reads the posts index
POST {{host}}/sites/{{site}}/sync
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.synced" == 2

# Scan: both fixture posts exceeded the 1-day freshness threshold
POST {{host}}/sites/{{site}}/scans/freshness
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 200
[Asserts]
jsonpath "$.detected" == 2

# Cold-start priority: crawl_hits is empty, so the queue order equals the
# posts date order, newest first — post-b (2026-06-03, epoch-day 20607)
# before post-a (2026-06-01, epoch-day 20605).
GET {{host}}/sites/{{site}}/queue?status=open
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 2
jsonpath "$.items[0].slug" == "post-b"
jsonpath "$.items[0].priority" == 20607
jsonpath "$.items[1].slug" == "post-a"
jsonpath "$.items[1].priority" == 20605
jsonpath "$.items[0].kind" == "refresh"
jsonpath "$.items[0].status" == "open"

# Re-scan: detection repeats but the insert is duplicate-free
POST {{host}}/sites/{{site}}/scans/freshness
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 200
[Asserts]
jsonpath "$.detected" == 2

# ... zero duplicate rows were enqueued
GET {{host}}/sites/{{site}}/queue?kind=refresh
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 2

# Export: both open items land in the bare repo as quests/queue/*.yaml
POST {{host}}/sites/{{site}}/queue/export
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.exported" == 2
jsonpath "$.consumed" == 0

# Both rows transitioned open → exported
GET {{host}}/sites/{{site}}/queue?status=open
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 0

GET {{host}}/sites/{{site}}/queue?status=exported
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 2

# Re-export: identical content is a no-op (no commit, nothing transitions)
POST {{host}}/sites/{{site}}/queue/export
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.exported" == 0
jsonpath "$.consumed" == 0
