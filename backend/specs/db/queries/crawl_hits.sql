-- name: CrawlHitAggSumsJson :one
-- Per-article crawl-hit sums as a JSON text scalar — the freshness scanner's
-- priority signal. SUM(hits + md_hits): .md consumption is direct agent
-- traffic, so it weighs in alongside page hits (Phase012). Empty set must be
-- '[]', not NULL (cold start: the scorer falls back to date recency).
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'lang', s.lang,
           'section', s.section,
           'slug', s.slug,
           'hits', s.hits
       ) ORDER BY s.lang, s.section, s.slug), '[]'::jsonb)::text
FROM (
    SELECT lang, section, slug, SUM(hits + md_hits)::BIGINT AS hits
    FROM crawl_hits
    WHERE site_id = @site_id
    GROUP BY lang, section, slug
) AS s;

-- name: CrawlHitUpsertFromJson :exec
-- Batch upsert of one ingest's aggregated rows. ON CONFLICT adds (never
-- replaces): the unique key (site_id, hit_date, bot, lang, section, slug) is
-- immutable and a later ingest of the same UTC date accumulates onto it.
-- Zero duplicate accumulation across re-ingests is the cursor's guarantee
-- (closed-hour boundaries), not this query's.
WITH incoming AS (
    SELECT (e->>'hit_date')::DATE AS hit_date,
           (e->>'bot')::TEXT      AS bot,
           (e->>'lang')::TEXT     AS lang,
           (e->>'section')::TEXT  AS section,
           (e->>'slug')::TEXT     AS slug,
           (e->>'hits')::BIGINT   AS hits,
           (e->>'md_hits')::BIGINT AS md_hits
    FROM jsonb_array_elements(@hits_json::jsonb) AS e
)
INSERT INTO crawl_hits (site_id, hit_date, bot, lang, section, slug, hits, md_hits)
SELECT @site_id, i.hit_date, i.bot, i.lang, i.section, i.slug, i.hits, i.md_hits
FROM incoming i
ON CONFLICT (site_id, hit_date, bot, lang, section, slug) DO UPDATE
SET hits = crawl_hits.hits + EXCLUDED.hits,
    md_hits = crawl_hits.md_hits + EXCLUDED.md_hits;

-- name: CrawlHitListFiltered :many
-- +no-pagination
-- Report/scanner lookup; empty-string filters mean "no filter". from/to
-- bound hit_date inclusively (UTC dates).
SELECT * FROM crawl_hits
WHERE site_id = @site_id
  AND (@lang::text = '' OR lang = @lang::text)
  AND (@section::text = '' OR section = @section::text)
  AND (@slug::text = '' OR slug = @slug::text)
  AND (@from_date::text = '' OR hit_date >= (@from_date::text)::date)
  AND (@to_date::text = '' OR hit_date <= (@to_date::text)::date)
ORDER BY hit_date, bot, lang, section, slug;

-- name: CrawlHitAggBotMonthJson :one
-- Per-(bot, article) sums over the report window: the 30 days ending on the
-- last day of @ym, both ends included ('' = the last closed month, UTC —
-- the only branch where NOW() enters; an explicit ym is clock-free). The
-- bot name stays unclassified: the training/search/fetch split lives only
-- in the pkg/bots Go dictionary, so SQL must never attempt it.
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            + interval '1 month - 1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'bot', s.bot,
           'lang', s.lang,
           'section', s.section,
           'slug', s.slug,
           'hits', s.hits,
           'md_hits', s.md_hits
       ) ORDER BY s.bot, s.lang, s.section, s.slug), '[]'::jsonb)::text
FROM (
    SELECT c.bot, c.lang, c.section, c.slug,
           SUM(c.hits)::BIGINT AS hits, SUM(c.md_hits)::BIGINT AS md_hits
    FROM crawl_hits c, bounds b
    WHERE c.site_id = @site_id
      AND c.hit_date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY c.bot, c.lang, c.section, c.slug
) AS s;

-- name: CrawlHitAggBotPrevMonthJson :one
-- The same aggregate over the previous month's window (month-over-month
-- trend): the 30 days ending on the last day of the month before @ym. The
-- month shift happens here because SSaC is linear — no ym arithmetic can
-- run between sequence steps.
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            - interval '1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'bot', s.bot,
           'lang', s.lang,
           'section', s.section,
           'slug', s.slug,
           'hits', s.hits,
           'md_hits', s.md_hits
       ) ORDER BY s.bot, s.lang, s.section, s.slug), '[]'::jsonb)::text
FROM (
    SELECT c.bot, c.lang, c.section, c.slug,
           SUM(c.hits)::BIGINT AS hits, SUM(c.md_hits)::BIGINT AS md_hits
    FROM crawl_hits c, bounds b
    WHERE c.site_id = @site_id
      AND c.hit_date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY c.bot, c.lang, c.section, c.slug
) AS s;
