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

# Operator-only: aggregate CloudFront logs into crawl_hits (daily cron)
allow if {
    input.action == "IngestCrawl"
    input.resource == "crawl_hit"
    input.claims.role == "operator"
}

# Operator-only: list aggregated crawl hits (reports and scanners)
allow if {
    input.action == "ListCrawlHits"
    input.resource == "crawl_hit"
    input.claims.role == "operator"
}

# Operator-only: collect Search Analytics rows into gsc_snapshots (daily cron)
allow if {
    input.action == "IngestGSC"
    input.resource == "gsc_snapshot"
    input.claims.role == "operator"
}

# Operator-only: register a citation-sampling query
allow if {
    input.action == "CreateCitationQuery"
    input.resource == "citation_query"
    input.claims.role == "operator"
}

# Operator-only: list citation-sampling queries
allow if {
    input.action == "ListCitationQueries"
    input.resource == "citation_query"
    input.claims.role == "operator"
}

# Operator-only: deactivate a citation-sampling query (soft delete)
allow if {
    input.action == "DeleteCitationQuery"
    input.resource == "citation_query"
    input.claims.role == "operator"
}

# Operator-only: run one citation-sampling round (weekly cron)
allow if {
    input.action == "SampleCitations"
    input.resource == "citation_sample"
    input.claims.role == "operator"
}

# Operator-only: read the citation sample time series
allow if {
    input.action == "ListCitations"
    input.resource == "citation_sample"
    input.claims.role == "operator"
}

# Operator-only: assemble and publish the monthly visibility report
allow if {
    input.action == "GenerateMonthlyReport"
    input.resource == "report"
    input.claims.role == "operator"
}

# Operator-only: read a stored monthly visibility report
allow if {
    input.action == "GetMonthlyReport"
    input.resource == "report"
    input.claims.role == "operator"
}
