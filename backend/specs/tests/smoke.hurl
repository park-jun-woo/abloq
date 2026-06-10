# Login rejects a wrong password
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "operator@abloq.test",
  "password": "wrong-password"
}
HTTP 401

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

# Sync — parse the fixture blog repository (2 posts) into the index
POST {{host}}/sync
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.synced" == 2

# ListPosts — the index holds exactly the 2 fixture posts
GET {{host}}/posts
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.posts" count == 2

# HookDeployed — deploy webhook records 3 pending receipts (1 URL × 3 kinds)
POST {{host}}/hooks/deployed
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "smoke-1",
  "changed": ["https://fixture.example.com/tech/post-a/"]
}
HTTP 202
[Asserts]
jsonpath "$.planned" == 3

# ProcessArchive — execute the pending receipts against the archive stub
POST {{host}}/archive/process
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "limit": 100
}
HTTP 200
[Asserts]
jsonpath "$.processed" >= 3

# RetryReceipts — nothing failed in smoke; rearm is a no-op
POST {{host}}/receipts/retry
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "smoke-1",
  "kind_filter": ""
}
HTTP 200
[Asserts]
jsonpath "$.pending" >= 0

# ListReceipts — all 3 smoke receipts are done
GET {{host}}/receipts?deploy_id=smoke-1&status=done
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.receipts" count == 3
