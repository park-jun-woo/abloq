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
