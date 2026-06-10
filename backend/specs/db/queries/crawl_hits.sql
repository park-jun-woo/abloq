-- name: CrawlHitAggSumsJson :one
-- Per-article crawl-hit sums as a JSON text scalar — the freshness scanner's
-- priority signal. Empty set must be '[]', not NULL (cold start: the
-- scorer falls back to date recency).
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'lang', s.lang,
           'section', s.section,
           'slug', s.slug,
           'hits', s.hits
       ) ORDER BY s.lang, s.section, s.slug), '[]'::jsonb)::text
FROM (
    SELECT lang, section, slug, SUM(hits)::BIGINT AS hits
    FROM crawl_hits
    GROUP BY lang, section, slug
) AS s;
