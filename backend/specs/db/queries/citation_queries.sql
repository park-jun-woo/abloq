-- name: CitationQueryCreate :one
INSERT INTO citation_queries (site_id, lang, section, slug, query_text)
VALUES (@site_id, @lang, @section, @slug, @query_text)
RETURNING *;

-- name: CitationQueryFindByID :one
-- Site-scoped lookup: another site's query id must read as "not found" —
-- the soft-delete endpoint would otherwise reach across sites.
SELECT * FROM citation_queries WHERE site_id = @site_id AND id = @id;

-- name: CitationQueryListFiltered :many
-- +no-pagination
-- Operational lookup; empty-string filters mean "no filter". active rides
-- as text ('', 'true', 'false') so one query serves all three states.
SELECT * FROM citation_queries
WHERE site_id = @site_id
  AND (@slug::text = '' OR slug = @slug::text)
  AND (@active_filter::text = '' OR active = (@active_filter::text)::boolean)
ORDER BY id;

-- name: CitationQuerySoftDelete :exec
-- Soft delete only: citation_samples references this row by FK — a hard
-- DELETE would destroy the recorded time series. The runner skips
-- active=false queries; the history stays queryable.
UPDATE citation_queries SET active = FALSE WHERE site_id = @site_id AND id = @id;

-- name: CitationQueryAggActiveJson :one
-- The active query set as a JSON text scalar for the sampling func (the
-- func never sees db model types — IngestCrawl precedent). id order is the
-- budget-cut order. Empty set must be '[]', not NULL.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'id', id,
           'query_text', query_text
       ) ORDER BY id), '[]'::jsonb)::text
FROM citation_queries
WHERE site_id = @site_id AND active = TRUE;
