# GSC polling + citation sampling scenario (Phase013). Assumes fresh
# gsc_snapshots/citation_queries/citation_samples tables and an abloqd
# started against the archive stub with GSC_SEARCH_API_BASE/GSC_TOKEN_URL,
# GSC_LOOKBACK_DAYS=3, GSC_INSPECT_RECENT_DAYS=36500 (both fixture posts
# always qualify), the three engine keys + base URLs pointed at the stub,
# and CITATION_INTERVAL_MS=0. The shared fixture blog.yaml sets
# citation_budget: 2. Runs before smoke (smoke re-runs every op loosely).

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

# Ingest 1: empty gsc_snapshots — '' cursor covers GSC_LOOKBACK_DAYS=3
# closed days, the stub answers 2 pages per day.
POST {{host}}/sites/{{site}}/ingest/gsc
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "inspect": false
}
HTTP 200
[Asserts]
jsonpath "$.days" == 3
jsonpath "$.rows" == 6
jsonpath "$.inspected" == 0
jsonpath "$.inspections" == "[]"

# Ingest 2: the MAX(snap_date) cursor sits at the last closed day —
# idempotent re-run collects nothing (zero duplicate accumulation).
POST {{host}}/sites/{{site}}/ingest/gsc
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "inspect": false
}
HTTP 200
[Asserts]
jsonpath "$.days" == 0
jsonpath "$.rows" == 0

# Ingest 3: opt-in URL Inspection — both fixture posts are recent under the
# huge GSC_INSPECT_RECENT_DAYS window; the stub verdict is PASS. Nothing is
# stored: the summary only rides in the response.
POST {{host}}/sites/{{site}}/ingest/gsc
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "inspect": true
}
HTTP 200
[Asserts]
jsonpath "$.days" == 0
jsonpath "$.rows" == 0
jsonpath "$.inspected" == 2
jsonpath "$.inspections" contains "PASS"

# Register the standard query set: one per fixture post.
POST {{host}}/sites/{{site}}/citation-queries
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "lang": "ko",
  "section": "tech",
  "slug": "post-a",
  "query_text": "best agentic blog framework"
}
HTTP 201
[Asserts]
jsonpath "$.slug" == "post-a"
jsonpath "$.active" == true

POST {{host}}/sites/{{site}}/citation-queries
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "lang": "ko",
  "section": "tech",
  "slug": "post-b",
  "query_text": "static blog geo checklist"
}
HTTP 201
[Captures]
qid_b: jsonpath "$.id"

# Both queries are listed; the slug filter narrows to one.
GET {{host}}/sites/{{site}}/citation-queries
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.queries" count == 2

GET {{host}}/sites/{{site}}/citation-queries?slug=post-a
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.queries" count == 1
jsonpath "$.queries[0].query_text" == "best agentic blog framework"

# Sampling round 1: 3 stub engines x 2 active queries (budget 2 covers
# both). Every stub answer cites the fixture domain — all rows cited.
POST {{host}}/sites/{{site}}/sample/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.engines" == 3
jsonpath "$.sampled" == 6

# The time series: engine-major insertion order (perplexity, openai,
# anthropic), extractor v1, evidence JSON carries the matched URL.
GET {{host}}/sites/{{site}}/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.samples" count == 6
jsonpath "$.samples[0].engine" == "perplexity"
jsonpath "$.samples[0].cited" == true
jsonpath "$.samples[0].extractor_version" == "v1"
jsonpath "$.samples[0].evidence" contains "fixture.example.com"
jsonpath "$.samples[2].engine" == "openai"
jsonpath "$.samples[4].engine" == "anthropic"
jsonpath "$.samples[4].cited" == true

# Slug filter: only the post-a query's samples (one per engine).
GET {{host}}/sites/{{site}}/citations?slug=post-a
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.samples" count == 3
jsonpath "$.samples[0].cited" == true

# Soft delete post-b's query: active flips, the row (and its samples) stay.
DELETE {{host}}/sites/{{site}}/citation-queries/{{qid_b}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.active" == false

GET {{host}}/sites/{{site}}/citation-queries?active_filter=true
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.queries" count == 1
jsonpath "$.queries[0].slug" == "post-a"

# ... and the recorded post-b samples survive the soft delete (FK history).
GET {{host}}/sites/{{site}}/citations?slug=post-b
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.samples" count == 3

# Deleting a query that does not exist is a 404 (guard path).
DELETE {{host}}/sites/{{site}}/citation-queries/999999
Authorization: Bearer {{token}}
HTTP 404

# Sampling round 2: only the surviving active query runs (3 engines x 1).
POST {{host}}/sites/{{site}}/sample/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.engines" == 3
jsonpath "$.sampled" == 3

GET {{host}}/sites/{{site}}/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.samples" count == 9
