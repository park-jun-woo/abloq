# Smoke — covers every operationId once. Runtime env requirements:
#   - archive stub (backend/scripts/archive-stub) + WAYBACK_/INDEXNOW_/GSC_*
#     overrides (Phase008),
#   - QUEUE_EXPORT_REPO_URL/QUEUE_EXPORT_WORKDIR pointing at a local bare
#     repo (backend/scripts/queue-export-test/setup-bare.sh) — the export
#     step pushes for real (Phase009),
#   - CF_LOG_SOURCE pointing at backend/fixtures/cflogs (Phase012).

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

# ScanFreshness — both fixture posts exceed the 1-day threshold; the insert
# itself is duplicate-free, so re-runs detect 2 and enqueue at most 2
POST {{host}}/scans/freshness
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 200
[Asserts]
jsonpath "$.detected" == 2

# ListQueue — the refresh candidates are in the queue (any status: the
# freshness scenario may already have exported them in a shared-DB run)
GET {{host}}/queue?kind=refresh
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 2

# ScanEvidence — post-a's unsourced claims enqueue 1 evidence item; post-b's
# dead link needs 3 consecutive failed scans before rot confirms (needs the
# archive stub on LINKCHECK_HOST_OVERRIDE: /dead-* paths answer 404)
POST {{host}}/scans/evidence
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" >= 1

# ScanCluster — coverage on the shared fixture is loose (the exact-count
# oracle lives in scenario-cluster.hurl against the dedicated cluster-blog
# fixture on its own abloqd instance — BLOG_REPO_PATH is process-wide)
POST {{host}}/scans/cluster
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" >= 0

# ExportQueue — pushes open items to the bare repo (QUEUE_EXPORT_*);
# a no-op when everything is already exported
POST {{host}}/queue/export
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.exported" >= 0
jsonpath "$.consumed" >= 0

# IngestCrawl — re-run after the crawl scenario already advanced the cursor
# (CF_LOG_SOURCE fixture logs): a no-op on the shared DB, a full ingest on a
# fresh one (the exact-count oracle lives in scenario-crawl.hurl)
POST {{host}}/ingest/crawl
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.files" >= 0
jsonpath "$.hits" >= 0

# ListCrawlHits — the fixture log hits are aggregated per article
GET {{host}}/crawl-hits
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.hits" count >= 0

# IngestGSC — re-run after the gsc-citation scenario already advanced the
# MAX(snap_date) cursor: a no-op on the shared DB, a full lookback on a
# fresh one (the exact-count oracle lives in scenario-gsc-citation.hurl)
POST {{host}}/ingest/gsc
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "inspect": false
}
HTTP 200
[Asserts]
jsonpath "$.days" >= 0
jsonpath "$.rows" >= 0

# CreateCitationQuery — register a smoke query (soft-deleted right below)
POST {{host}}/citation-queries
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "lang": "ko",
  "section": "tech",
  "slug": "post-a",
  "query_text": "smoke probe query"
}
HTTP 201
[Captures]
smoke_qid: jsonpath "$.id"
[Asserts]
jsonpath "$.active" == true

# ListCitationQueries — at least the smoke query is listed
GET {{host}}/citation-queries
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.queries" count >= 1

# SampleCitations — engine keys and budget come from the runtime env; on
# the stub wiring this records rows, elsewhere it may be a clean no-op
POST {{host}}/sample/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.engines" >= 0
jsonpath "$.sampled" >= 0

# ListCitations — the recorded time series is readable
GET {{host}}/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.samples" count >= 0

# DeleteCitationQuery — soft delete the smoke query (history preserved)
DELETE {{host}}/citation-queries/{{smoke_qid}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.active" == false

# GenerateMonthlyReport — default ym ('' = the last closed month); queue and
# citation layers of that window are deterministically empty (created_at /
# run_at are NOW(), never inside a closed month)
POST {{host}}/reports/monthly
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 200
[Captures]
report_ym: jsonpath "$.ym"
[Asserts]
jsonpath "$.ym" matches "^[0-9]{4}-[0-9]{2}$"
jsonpath "$.articles" >= 2

# GetMonthlyReport — the stored row is the lookup truth
GET {{host}}/reports/monthly/{{report_ym}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.ym" == "{{report_ym}}"
jsonpath "$.markdown" contains "Visibility report"
jsonpath "$.report_json" contains "\"articles\""
