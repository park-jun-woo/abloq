package authz

default allow := false

# Operator-only: sync the posts index from the blog repository
allow if {
    input.action == "Sync"
    input.resource == "post"
    input.claims.role == "operator"
}

# Operator-only: list the indexed posts
allow if {
    input.action == "ListPosts"
    input.resource == "post"
    input.claims.role == "operator"
}
