-- name: CitationCheckAggAllJson :one
-- Previous probe state as a JSON text scalar for the evidence scanner func
-- (the func never sees db model types). Empty set must be '[]', not NULL.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'url', url,
           'lang', lang,
           'section', section,
           'slug', slug,
           'status', status,
           'consecutive_failures', consecutive_failures
       ) ORDER BY id), '[]'::jsonb)::text
FROM citation_checks
WHERE site_id = @site_id;

-- name: CitationCheckUpsertFromJson :exec
-- Batch upsert of this scan's probe results. The scanner emits one entry per
-- (url, lang, section, slug) key — within one article duplicate URLs are
-- collapsed upstream, so the batch never touches a key twice. consecutive
-- failure counting (fail: prev+1, ok: reset to 0) happens in the func's pkg;
-- this query only persists the computed state.
WITH incoming AS (
    SELECT (e->>'url')::TEXT                  AS url,
           (e->>'lang')::TEXT                 AS lang,
           (e->>'section')::TEXT              AS section,
           (e->>'slug')::TEXT                 AS slug,
           (e->>'status')::TEXT               AS status,
           (e->>'consecutive_failures')::BIGINT AS consecutive_failures
    FROM jsonb_array_elements(@checks_json::jsonb) AS e
)
INSERT INTO citation_checks (site_id, url, lang, section, slug, status, consecutive_failures)
SELECT @site_id, i.url, i.lang, i.section, i.slug, i.status, i.consecutive_failures
FROM incoming i
ON CONFLICT (site_id, url, lang, section, slug) DO UPDATE
SET status = EXCLUDED.status,
    consecutive_failures = EXCLUDED.consecutive_failures,
    last_checked_at = NOW(),
    updated_at = NOW();
