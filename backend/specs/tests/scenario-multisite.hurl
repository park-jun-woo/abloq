# Multisite isolation scenario (Phase020) — one abloqd instance, the
# sites.yaml fixture declaring three sites: `default` (the shared fixture
# blog, already exercised by the earlier scenarios), `beta` (a copy of the
# same fixture — same slugs on purpose, the coexistence oracle) and
# `retired` (declared active: false — the deactivated-site 404 oracle).
# Runs late in the shared batch (after scenario-report, before smoke):
# `default` has synced posts, queue rows and the seeded 2026-04 report,
# while `beta` is untouched — exactly the precondition the isolation
# asserts need. Covers the six Phase020 isolation oracles plus the
# unregistered-site 404 of every site-scoped operation (XOH-12/13).

# Login as the operator (tokens are instance-global — sites share operators)
POST {{host}}/auth/login
Content-Type: application/json
[Options]
variable: site="default"
variable: beta="beta"
variable: ghost="nosuch"
variable: retired="retired"
variable: ms_ym="2026-04"
{
  "email": "operator@abloq.test",
  "password": "abloq-operator-test"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"

# The registry lists the active sites; retired only shows unfiltered.
GET {{host}}/sites?active_filter=true
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.sites" count == 2
jsonpath "$.sites[*].name" includes "default"
jsonpath "$.sites[*].name" includes "beta"
jsonpath "$.sites[*].name" not includes "retired"

GET {{host}}/sites
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.sites" count == 3
jsonpath "$.sites[*].name" includes "retired"

GET {{host}}/sites?active_filter=false
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.sites" count == 1
jsonpath "$.sites[0].name" == "retired"
jsonpath "$.sites[0].active" == false

# ① A sync ≠ B rows: default synced long ago, beta never — beta reads empty.
GET {{host}}/sites/{{beta}}/posts
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.posts" count == 0

# ② A's queue is invisible through B: default holds refresh+evidence rows
# from the earlier scenarios, beta's queue is empty.
GET {{host}}/sites/{{site}}/queue
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count >= 2

GET {{host}}/sites/{{beta}}/queue
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 0

# ③ Same slugs coexist: beta syncs the same fixture content — both sites
# now hold (lang=ko, section=tech, slug=post-a|post-b) without collision.
POST {{host}}/sites/{{beta}}/sync
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.synced" == 2

GET {{host}}/sites/{{beta}}/posts
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.posts" count == 2
jsonpath "$.posts[0].slug" == "post-a"
jsonpath "$.posts[1].slug" == "post-b"

GET {{host}}/sites/{{site}}/posts
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.posts" count == 2
jsonpath "$.posts[0].slug" == "post-a"

# ②' ... and the reverse direction: beta's freshness scan enqueues its own
# 2 open rows; default's refresh rows stay exported (the earlier export) —
# neither side sees the other's items.
POST {{host}}/sites/{{beta}}/scans/freshness
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 200
[Asserts]
jsonpath "$.detected" == 2

GET {{host}}/sites/{{beta}}/queue?kind=refresh&status=open
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 2

GET {{host}}/sites/{{site}}/queue?kind=refresh&status=open
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 0

# ④ (site_id, ym) report independence: default's seeded 2026-04 report
# exists, beta's does not — until beta generates its own, which never
# inherits default's seeded crawl/GSC numbers.
GET {{host}}/sites/{{beta}}/reports/monthly/{{ms_ym}}
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{beta}}/reports/monthly
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": "2026-04"
}
HTTP 200
[Asserts]
jsonpath "$.ym" == "2026-04"
jsonpath "$.articles" == 2
jsonpath "$.published" == true

GET {{host}}/sites/{{beta}}/reports/monthly/{{ms_ym}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.markdown" contains "# Visibility report 2026-04"
jsonpath "$.markdown" not contains "| ko/tech/post-a | 2026-06-01 | 7 | 0 | 3 | 2 | 120 | 8 | 0 | 136 |"

GET {{host}}/sites/{{site}}/reports/monthly/{{ms_ym}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.markdown" contains "| ko/tech/post-a | 2026-06-01 | 7 | 0 | 3 | 2 | 120 | 8 | 0 | 136 |"

# ⑤ Unregistered site is a plain 404 on every site-scoped operation
# (path-scoped fail-closed: a missing site can never read or write).
POST {{host}}/sites/{{ghost}}/sync
Authorization: Bearer {{token}}
HTTP 404

GET {{host}}/sites/{{ghost}}/posts
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/hooks/deployed
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "ms-404",
  "changed": ["https://fixture.example.com/tech/post-a/"]
}
HTTP 404

POST {{host}}/sites/{{ghost}}/archive/process
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "limit": 1
}
HTTP 404

POST {{host}}/sites/{{ghost}}/receipts/retry
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "deploy_id": "",
  "kind_filter": ""
}
HTTP 404

GET {{host}}/sites/{{ghost}}/receipts
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/scans/freshness
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 404

POST {{host}}/sites/{{ghost}}/scans/evidence
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/scans/cluster
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/queue/export
Authorization: Bearer {{token}}
HTTP 404

GET {{host}}/sites/{{ghost}}/queue
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/ingest/crawl
Authorization: Bearer {{token}}
HTTP 404

GET {{host}}/sites/{{ghost}}/crawl-hits
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/ingest/gsc
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "inspect": false
}
HTTP 404

POST {{host}}/sites/{{ghost}}/citation-queries
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "lang": "ko",
  "section": "tech",
  "slug": "post-a",
  "query_text": "unregistered probe"
}
HTTP 404

GET {{host}}/sites/{{ghost}}/citation-queries
Authorization: Bearer {{token}}
HTTP 404

DELETE {{host}}/sites/{{ghost}}/citation-queries/1
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/sample/citations
Authorization: Bearer {{token}}
HTTP 404

GET {{host}}/sites/{{ghost}}/citations
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{ghost}}/reports/monthly
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "ym": ""
}
HTTP 404

GET {{host}}/sites/{{ghost}}/reports/monthly/{{ms_ym}}
Authorization: Bearer {{token}}
HTTP 404

# ⑥ A deactivated site answers the same 404 — declared rows survive in the
# DB (FK history) but stop serving the moment sites.yaml drops them.
GET {{host}}/sites/{{retired}}/posts
Authorization: Bearer {{token}}
HTTP 404

POST {{host}}/sites/{{retired}}/sync
Authorization: Bearer {{token}}
HTTP 404
