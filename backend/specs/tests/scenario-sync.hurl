# Sync scenario — fixture repository (2 posts) lands as 2 indexed rows
# with parsed front matter fields and the blog.yaml URL contract.

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

# Sync the fixture blog repository
POST {{host}}/sites/{{site}}/sync
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.synced" == 2

# Sync is idempotent — a second run upserts the same 2 rows
POST {{host}}/sites/{{site}}/sync
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.synced" == 2

# The index exposes the parsed front matter and computed metrics
GET {{host}}/sites/{{site}}/posts
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.posts" count == 2
jsonpath "$.posts[0].lang" == "ko"
jsonpath "$.posts[0].section" == "tech"
jsonpath "$.posts[0].slug" == "post-a"
jsonpath "$.posts[0].title" == "포스트 A"
jsonpath "$.posts[0].date" == "2026-06-01"
jsonpath "$.posts[0].lastmod" == "2026-06-05"
jsonpath "$.posts[0].url" == "https://fixture.example.com/tech/post-a/"
jsonpath "$.posts[0].word_count" > 0
jsonpath "$.posts[0].internal_links" == 1
jsonpath "$.posts[0].source_count" == 2
jsonpath "$.posts[1].slug" == "post-b"
jsonpath "$.posts[1].url" == "https://fixture.example.com/tech/post-b/"
