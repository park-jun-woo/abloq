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
