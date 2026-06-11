# Authorization scenario — every endpoint is operator-only.
# A non-operator (viewer) authenticates but is denied by the Rego policy.

# Login as the viewer
POST {{host}}/auth/login
Content-Type: application/json
[Options]
variable: site="default"
{
  "email": "viewer@abloq.test",
  "password": "abloq-viewer-test"
}
HTTP 200
[Captures]
viewer_token: jsonpath "$.access_token"

# ListSites is forbidden for non-operators
GET {{host}}/sites
Authorization: Bearer {{viewer_token}}
HTTP 403

# Sync is forbidden for non-operators
POST {{host}}/sites/{{site}}/sync
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListPosts is forbidden for non-operators
GET {{host}}/sites/{{site}}/posts
Authorization: Bearer {{viewer_token}}
HTTP 403

# HookDeployed is forbidden for non-operators
POST {{host}}/sites/{{site}}/hooks/deployed
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "deploy_id": "forbidden-1",
  "changed": ["https://fixture.example.com/tech/post-a/"]
}
HTTP 403

# ProcessArchive is forbidden for non-operators
POST {{host}}/sites/{{site}}/archive/process
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "limit": 1
}
HTTP 403

# RetryReceipts is forbidden for non-operators
POST {{host}}/sites/{{site}}/receipts/retry
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "deploy_id": "",
  "kind_filter": ""
}
HTTP 403

# ListReceipts is forbidden for non-operators
GET {{host}}/sites/{{site}}/receipts
Authorization: Bearer {{viewer_token}}
HTTP 403

# ScanFreshness is forbidden for non-operators
POST {{host}}/sites/{{site}}/scans/freshness
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 403

# ScanEvidence is forbidden for non-operators
POST {{host}}/sites/{{site}}/scans/evidence
Authorization: Bearer {{viewer_token}}
HTTP 403

# ScanCluster is forbidden for non-operators
POST {{host}}/sites/{{site}}/scans/cluster
Authorization: Bearer {{viewer_token}}
HTTP 403

# ExportQueue is forbidden for non-operators
POST {{host}}/sites/{{site}}/queue/export
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListQueue is forbidden for non-operators
GET {{host}}/sites/{{site}}/queue
Authorization: Bearer {{viewer_token}}
HTTP 403

# IngestCrawl is forbidden for non-operators
POST {{host}}/sites/{{site}}/ingest/crawl
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListCrawlHits is forbidden for non-operators
GET {{host}}/sites/{{site}}/crawl-hits
Authorization: Bearer {{viewer_token}}
HTTP 403

# IngestGSC is forbidden for non-operators
POST {{host}}/sites/{{site}}/ingest/gsc
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "inspect": false
}
HTTP 403

# CreateCitationQuery is forbidden for non-operators
POST {{host}}/sites/{{site}}/citation-queries
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "lang": "ko",
  "section": "tech",
  "slug": "post-a",
  "query_text": "forbidden probe"
}
HTTP 403

# ListCitationQueries is forbidden for non-operators
GET {{host}}/sites/{{site}}/citation-queries
Authorization: Bearer {{viewer_token}}
HTTP 403

# DeleteCitationQuery is forbidden for non-operators
DELETE {{host}}/sites/{{site}}/citation-queries/1
Authorization: Bearer {{viewer_token}}
HTTP 403

# SampleCitations is forbidden for non-operators
POST {{host}}/sites/{{site}}/sample/citations
Authorization: Bearer {{viewer_token}}
HTTP 403

# ListCitations is forbidden for non-operators
GET {{host}}/sites/{{site}}/citations
Authorization: Bearer {{viewer_token}}
HTTP 403

# GenerateMonthlyReport is forbidden for non-operators
POST {{host}}/sites/{{site}}/reports/monthly
Authorization: Bearer {{viewer_token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 403

# GetMonthlyReport is forbidden for non-operators
GET {{host}}/sites/{{site}}/reports/monthly/{{forbidden_ym}}
Authorization: Bearer {{viewer_token}}
[Options]
variable: forbidden_ym="2026-04"
HTTP 403
