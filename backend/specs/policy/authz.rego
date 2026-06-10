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

# Operator-only: deploy webhook records pending archive receipts
allow if {
    input.action == "HookDeployed"
    input.resource == "receipt"
    input.claims.role == "operator"
}

# Operator-only: execute pending archive receipts (external calls)
allow if {
    input.action == "ProcessArchive"
    input.resource == "receipt"
    input.claims.role == "operator"
}

# Operator-only: rearm failed/deferred receipts back to pending
allow if {
    input.action == "RetryReceipts"
    input.resource == "receipt"
    input.claims.role == "operator"
}

# Operator-only: list archive receipts (quest-gate lookup)
allow if {
    input.action == "ListReceipts"
    input.resource == "receipt"
    input.claims.role == "operator"
}

# Operator-only: detect stale articles and enqueue refresh candidates
allow if {
    input.action == "ScanFreshness"
    input.resource == "queue_item"
    input.claims.role == "operator"
}

# Operator-only: detect unsourced claims/link rot and enqueue evidence candidates
allow if {
    input.action == "ScanEvidence"
    input.resource == "queue_item"
    input.claims.role == "operator"
}

# Operator-only: detect cluster violations and enqueue curation candidates
allow if {
    input.action == "ScanCluster"
    input.resource == "queue_item"
    input.claims.role == "operator"
}

# Operator-only: export open queue items to the blog repository
allow if {
    input.action == "ExportQueue"
    input.resource == "queue_item"
    input.claims.role == "operator"
}

# Operator-only: list queue items (operational lookup)
allow if {
    input.action == "ListQueue"
    input.resource == "queue_item"
    input.claims.role == "operator"
}
