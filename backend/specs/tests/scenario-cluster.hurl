# Cluster scanner scenario — scan → enqueue (violating articles only, exact
# counts) → re-scan idempotence (zero duplicates) → export to the local bare
# repo. Runs against a SECOND abloqd instance: fresh database, its own port,
# BLOG_REPO_PATH=backend/fixtures/cluster-blog and dedicated QUEUE_EXPORT_*
# bare repo (BLOG_REPO_PATH is process-wide env, and the shared fixture must
# not grow articles — scenario-sync asserts synced==2). The candidates
# payload is @sensitive (never in the QueueItem response), so the
# candidate-existence leg lives in backend/scripts/cluster-scan-test/run.sh
# over the exported bare clone, plus the CLI diff equivalence.
#
# Fixture graph (ko default, section tech, taxonomy [geo, abloq, hugo, lonely],
# min_internal_links 2):
#   hub    — clean: 2 outlinks (thin, offtax), 4 inlinks, shared in-taxonomy tags
#   thin   — min-internal-links: 1 outlink (hub); priority 1+3 candidates = 4
#   orphan — no-orphan-tag (lonely, in-taxonomy but alone); 1+4 = 5
#   island — tag-taxonomy (rogue) + no-isolated-post (0 inlinks); 2+4 = 6
#   offtax — tag-taxonomy (rogue, shared with island → not an orphan); 1+3 = 4
# The listing is priority DESC: island(6) orphan(5) then offtax/thin (4
# each — their relative order depends on batch-insert id assignment, which
# the INSERT...SELECT anti-join does not guarantee, so the tie is asserted
# by membership, not position).

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

# Scan: exactly the 4 violating articles are detected and enqueued
POST {{host}}/sites/{{site}}/scans/cluster
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" == 4

# The queue holds one open kind=cluster item per violating article —
# priority = violations + candidates (never now-derived), hub absent
GET {{host}}/sites/{{site}}/queue?kind=cluster
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 4
jsonpath "$.items[0].slug" == "island"
jsonpath "$.items[0].priority" == 6
jsonpath "$.items[0].kind" == "cluster"
jsonpath "$.items[0].lang" == "ko"
jsonpath "$.items[0].status" == "open"
jsonpath "$.items[1].slug" == "orphan"
jsonpath "$.items[1].priority" == 5
jsonpath "$.items[2].priority" == 4
jsonpath "$.items[3].priority" == 4
jsonpath "$.items[*].slug" includes "offtax"
jsonpath "$.items[*].slug" includes "thin"
jsonpath "$.items[*].slug" not includes "hub"

# Re-scan: detection repeats but the insert is duplicate-free
POST {{host}}/sites/{{site}}/scans/cluster
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.detected" == 4

# ... zero duplicate rows were enqueued
GET {{host}}/sites/{{site}}/queue?kind=cluster
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 4

# Export: the 4 open items land in the bare repo as quests/queue/*.yaml —
# the script leg then checks every cluster file carries the candidates key
POST {{host}}/sites/{{site}}/queue/export
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.exported" == 4
jsonpath "$.consumed" == 0

# All rows transitioned open → exported
GET {{host}}/sites/{{site}}/queue?status=open
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" count == 0

# Re-export: identical content is a no-op
POST {{host}}/sites/{{site}}/queue/export
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.exported" == 0
jsonpath "$.consumed" == 0
