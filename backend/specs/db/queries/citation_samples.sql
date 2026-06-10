-- name: CitationSampleInsertFromJson :exec
-- Batch insert of one sampling round. Samples are pure time-series records
-- (no unique key, run_at = DB NOW()): every round appends. budget 0 sends
-- '[]' and inserts nothing — the no-op oracle.
WITH incoming AS (
    SELECT (e->>'citation_queries_id')::BIGINT AS citation_queries_id,
           (e->>'engine')::TEXT                AS engine,
           (e->>'cited')::BOOLEAN              AS cited,
           (e->>'evidence')::TEXT              AS evidence,
           (e->>'extractor_version')::TEXT     AS extractor_version
    FROM jsonb_array_elements(@samples_json::jsonb) AS e
)
INSERT INTO citation_samples (citation_queries_id, engine, cited, evidence, extractor_version)
SELECT i.citation_queries_id, i.engine, i.cited, i.evidence, i.extractor_version
FROM incoming i;

-- name: CitationSampleListFiltered :many
-- +no-pagination
-- Citation time series; the slug filter joins through citation_queries via
-- a subquery so the row shape stays the bare model (the response converter
-- is table-shaped). Empty-string slug means "no filter"; run_at then id is
-- the series order.
SELECT * FROM citation_samples
WHERE (@slug::text = ''
       OR citation_queries_id IN (SELECT id FROM citation_queries WHERE slug = @slug::text))
ORDER BY run_at, id;

-- name: CitationSampleAggMonthJson :one
-- Per-article cited/total sample counts over the report window ('' = the
-- last closed month, UTC; run_at compared as a UTC date). Joined through
-- citation_queries for the article key. Priority input + report table only
-- — never a gate (§6.3).
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            + interval '1 month - 1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'lang', s.lang,
           'section', s.section,
           'slug', s.slug,
           'cited', s.cited,
           'total', s.total
       ) ORDER BY s.lang, s.section, s.slug), '[]'::jsonb)::text
FROM (
    SELECT q.lang, q.section, q.slug,
           COUNT(*) FILTER (WHERE cs.cited)::BIGINT AS cited,
           COUNT(*)::BIGINT AS total
    FROM citation_samples cs
    JOIN citation_queries q ON q.id = cs.citation_queries_id
    CROSS JOIN bounds b
    WHERE (cs.run_at AT TIME ZONE 'utc')::date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY q.lang, q.section, q.slug
) AS s;

-- name: CitationSampleAggPrevMonthJson :one
-- The same aggregate over the previous month's window (trend column).
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            - interval '1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'lang', s.lang,
           'section', s.section,
           'slug', s.slug,
           'cited', s.cited,
           'total', s.total
       ) ORDER BY s.lang, s.section, s.slug), '[]'::jsonb)::text
FROM (
    SELECT q.lang, q.section, q.slug,
           COUNT(*) FILTER (WHERE cs.cited)::BIGINT AS cited,
           COUNT(*)::BIGINT AS total
    FROM citation_samples cs
    JOIN citation_queries q ON q.id = cs.citation_queries_id
    CROSS JOIN bounds b
    WHERE (cs.run_at AT TIME ZONE 'utc')::date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY q.lang, q.section, q.slug
) AS s;
