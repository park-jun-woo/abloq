# Evidence scanner scenario — claims enqueue immediately, link rot only after
# 3 consecutive failed scans (citation_checks state). Assumes a fresh database
# and an abloqd started with LINKCHECK_HOST_OVERRIDE pointing at the archive
# stub (GET/HEAD: paths containing "dead" answer 404, the rest 200), so the
# fixture citation URLs never touch the real network. Fixture layout: post-a
# carries 2 unsourced numeric claims (its citations are alive), post-b carries
# the dead link and no claims — separate articles, because claims and rot of
# one article share the queue's idempotent key.

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

# Scan 1: post-a's claims item is detected and enqueued. post-b's dead link
# is now at 1 consecutive failure — not rot yet.
POST {{host}}/scans/evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" == 1

# Only the claims item is in the queue: 2 claims → priority 2.
GET {{host}}/queue?kind=evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 1
jsonpath "$.items[0].slug" == "post-a"
jsonpath "$.items[0].kind" == "evidence"
jsonpath "$.items[0].status" == "open"
jsonpath "$.items[0].priority" == 2

# Scan 2: detection repeats but the insert is duplicate-free; the dead link
# reaches 2 consecutive failures — still not rot.
POST {{host}}/scans/evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" == 1

# ... zero duplicate rows, and post-b is still not queued
GET {{host}}/queue?kind=evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 1
jsonpath "$.items[0].slug" == "post-a"

# Scan 3: the dead link hits 3 consecutive failures — rot confirms and
# post-b's evidence item joins the queue.
POST {{host}}/scans/evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" == 2

GET {{host}}/queue?kind=evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 2
jsonpath "$.items[0].slug" == "post-a"
jsonpath "$.items[0].priority" == 2
jsonpath "$.items[1].slug" == "post-b"
jsonpath "$.items[1].priority" == 1

# Scan 4: fully idempotent — both items detected, zero new rows.
POST {{host}}/scans/evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" == 2

GET {{host}}/queue?kind=evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 2
