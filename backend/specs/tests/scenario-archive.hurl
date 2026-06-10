# Archiver scenario — receipt-table queue end to end against the local
# archive stub (backend/scripts/archive-stub): happy path, idempotent
# re-webhook, partial failure (stub 5xx on "boom" targets) and retry rearm.
# Assumes a fresh database (deploy ids dep-arch-* are unused).

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

# Webhook: 2 changed URLs → 6 pending receipts (2 URLs × 3 kinds), 202 fast
POST {{host}}/hooks/deployed
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "dep-arch-1",
  "changed": [
    "https://fixture.example.com/tech/post-a/",
    "https://fixture.example.com/tech/post-b/"
  ]
}
HTTP 202
[Asserts]
jsonpath "$.planned" == 6

# All 6 receipts are pending — no external call happened yet
GET {{host}}/receipts?deploy_id=dep-arch-1&status=pending
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 6

# Process: external calls (stub) — every receipt lands done
POST {{host}}/archive/process
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "limit": 100
}
HTTP 200
[Asserts]
jsonpath "$.processed" >= 6
jsonpath "$.failed" == 0
jsonpath "$.deferred" == 0

GET {{host}}/receipts?deploy_id=dep-arch-1&status=done
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 6
jsonpath "$.receipts[0].deploy_id" == "dep-arch-1"
jsonpath "$.receipts[0].status" == "done"

# Idempotent re-webhook: same deploy_id + same URLs → nothing new planned
POST {{host}}/hooks/deployed
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "dep-arch-1",
  "changed": [
    "https://fixture.example.com/tech/post-a/",
    "https://fixture.example.com/tech/post-b/"
  ]
}
HTTP 202
[Asserts]
jsonpath "$.planned" == 0

# ... and no receipt was duplicated or knocked back to pending
GET {{host}}/receipts?deploy_id=dep-arch-1
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 6

# Partial failure: the stub returns 5xx for wayback saves of "boom" targets
POST {{host}}/hooks/deployed
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "dep-arch-2",
  "changed": ["https://fixture.example.com/tech/boom-post/"]
}
HTTP 202
[Asserts]
jsonpath "$.planned" == 3

POST {{host}}/archive/process
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "limit": 100
}
HTTP 200
[Asserts]
jsonpath "$.failed" == 1

# wayback failed, indexnow + gsc_index done — failure is isolated per target
GET {{host}}/receipts?deploy_id=dep-arch-2&status=failed
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 1
jsonpath "$.receipts[0].kind" == "wayback"

GET {{host}}/receipts?deploy_id=dep-arch-2&status=done
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 2

# Retry rearms the failed wayback receipt back to pending
POST {{host}}/receipts/retry
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "dep-arch-2",
  "kind_filter": "wayback"
}
HTTP 200
[Asserts]
# global pending count — other hurl files may hold transient pending rows
# when the suite runs in parallel, so assert the floor and pin the exact
# row below via the deploy_id-filtered lookup
jsonpath "$.pending" >= 1

GET {{host}}/receipts?deploy_id=dep-arch-2&status=pending
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 1
jsonpath "$.receipts[0].kind" == "wayback"
