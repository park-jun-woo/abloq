# Authorization scenario — every endpoint is operator-only.
# A non-operator (viewer) authenticates but is denied by the Rego policy.

# Login as the viewer
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "viewer@abloq.test",
  "password": "abloq-viewer-test"
}
HTTP 200
[Captures]
viewer_token: jsonpath "$.access_token"

# Sync is forbidden for non-operators
POST {{host}}/sync
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListPosts is forbidden for non-operators
GET {{host}}/posts
Authorization: Bearer {{viewer_token}}
HTTP 403

# HookDeployed is forbidden for non-operators
POST {{host}}/hooks/deployed
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "deploy_id": "forbidden-1",
  "changed": ["https://fixture.example.com/tech/post-a/"]
}
HTTP 403

# ProcessArchive is forbidden for non-operators
POST {{host}}/archive/process
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "limit": 1
}
HTTP 403

# RetryReceipts is forbidden for non-operators
POST {{host}}/receipts/retry
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "deploy_id": "",
  "kind_filter": ""
}
HTTP 403

# ListReceipts is forbidden for non-operators
GET {{host}}/receipts
Authorization: Bearer {{viewer_token}}
HTTP 403

# ScanFreshness is forbidden for non-operators
POST {{host}}/scans/freshness
Authorization: Bearer {{viewer_token}}
HTTP 403

# ScanEvidence is forbidden for non-operators
POST {{host}}/scans/evidence
Authorization: Bearer {{viewer_token}}
HTTP 403

# ScanCluster is forbidden for non-operators
POST {{host}}/scans/cluster
Authorization: Bearer {{viewer_token}}
HTTP 403

# ExportQueue is forbidden for non-operators
POST {{host}}/queue/export
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListQueue is forbidden for non-operators
GET {{host}}/queue
Authorization: Bearer {{viewer_token}}
HTTP 403

# IngestCrawl is forbidden for non-operators
POST {{host}}/ingest/crawl
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListCrawlHits is forbidden for non-operators
GET {{host}}/crawl-hits
Authorization: Bearer {{viewer_token}}
HTTP 403
