# Crawl ingest scenario — CloudFront logs to crawl_hits. Assumes a fresh
# crawl_hits/ingest_cursors table set and an abloqd started with
# CF_LOG_SOURCE=backend/fixtures/cflogs (two closed-hour .gz logs against the
# shared fixture blog: GPTBot 2 page hits + 1 .md hit on post-a, ClaudeBot 1
# hit and ChatGPT-User 1 hit (304) on post-b; a human, a PetalBot unknown
# candidate and a PerplexityBot 404 must not land in crawl_hits). Runs after
# scenario-freshness (its cold-start priority assert needs crawl_hits empty)
# and before smoke.

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

# Ingest 1: both closed-hour fixture files are aggregated wholesale.
# hits = 5 article hits total (4 html + 1 md).
POST {{host}}/ingest/crawl
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.files" == 2
jsonpath "$.hits" == 5

# All rows, canonical order (hit_date, bot, lang, section, slug)
GET {{host}}/crawl-hits
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.hits" count == 3
jsonpath "$.hits[0].bot" == "ChatGPT-User"
jsonpath "$.hits[0].slug" == "post-b"
jsonpath "$.hits[0].hits" == 1
jsonpath "$.hits[0].md_hits" == 0
jsonpath "$.hits[1].bot" == "ClaudeBot"
jsonpath "$.hits[1].slug" == "post-b"
jsonpath "$.hits[1].hits" == 1
jsonpath "$.hits[2].bot" == "GPTBot"
jsonpath "$.hits[2].slug" == "post-a"
jsonpath "$.hits[2].lang" == "ko"
jsonpath "$.hits[2].section" == "tech"
jsonpath "$.hits[2].hits" == 2
jsonpath "$.hits[2].md_hits" == 1

# Filtered lookup: post-a only — the .md twin accumulated as md_hits
GET {{host}}/crawl-hits?lang=ko&section=tech&slug=post-a
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.hits" count == 1
jsonpath "$.hits[0].hits" == 2
jsonpath "$.hits[0].md_hits" == 1

# Date-range filters: the fixture hits are all on 2026-06-01 (UTC)
GET {{host}}/crawl-hits?from_date=2026-06-02
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.hits" count == 0

GET {{host}}/crawl-hits?from_date=2026-06-01&to_date=2026-06-01
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.hits" count == 3

# Ingest 2: the cursor advanced past both closed hours — nothing to ingest.
POST {{host}}/ingest/crawl
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.files" == 0
jsonpath "$.hits" == 0

# ... and zero duplicate accumulation: every count is unchanged.
GET {{host}}/crawl-hits
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.hits" count == 3
jsonpath "$.hits[0].hits" == 1
jsonpath "$.hits[1].hits" == 1
jsonpath "$.hits[2].hits" == 2
jsonpath "$.hits[2].md_hits" == 1
