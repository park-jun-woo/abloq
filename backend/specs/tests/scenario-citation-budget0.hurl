# Citation sampling budget=0 no-op scenario (Phase013). Runs on the
# dedicated cluster instance (cluster-blog fixture): its blog.yaml leaves
# geo.citation_budget at the 0 default, so sampling is disabled even though
# the engine keys are set and an active query exists — the assert isolates
# the budget as the cause. (The recording oracle lives in
# scenario-gsc-citation.hurl on the shared instance — one blog.yaml cannot
# satisfy both.)

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

# An active query exists...
POST {{host}}/sites/{{site}}/citation-queries
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "lang": "ko",
  "section": "tech",
  "slug": "hub",
  "query_text": "topical authority hub article"
}
HTTP 201
[Asserts]
jsonpath "$.active" == true

# ... and the engines are wired, but budget 0 disables the round.
POST {{host}}/sites/{{site}}/sample/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.engines" == 3
jsonpath "$.sampled" == 0

# Nothing was recorded.
GET {{host}}/sites/{{site}}/citations
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.samples" count == 0
