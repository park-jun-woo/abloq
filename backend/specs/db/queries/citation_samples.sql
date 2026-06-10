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
