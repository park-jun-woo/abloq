-- name: IngestCursorAggAllJson :one
-- Cursor state as a JSON text scalar for the ingest func (the func never
-- sees db model types — ScanEvidence precedent). Empty set must be '[]'.
-- The cursor key is (site_id, source) — each site's ingest progresses
-- independently even when log sources share a bucket prefix style.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'source', source,
           'cursor_hour', cursor_hour
       ) ORDER BY source), '[]'::jsonb)::text
FROM ingest_cursors
WHERE site_id = @site_id;

-- name: IngestCursorUpsertFromJson :exec
-- Advance cursors to the closed-hour boundaries the ingest just covered.
WITH incoming AS (
    SELECT (e->>'source')::TEXT      AS source,
           (e->>'cursor_hour')::TEXT AS cursor_hour
    FROM jsonb_array_elements(@cursors_json::jsonb) AS e
)
INSERT INTO ingest_cursors (site_id, source, cursor_hour)
SELECT @site_id, i.source, i.cursor_hour
FROM incoming i
ON CONFLICT (site_id, source) DO UPDATE
SET cursor_hour = EXCLUDED.cursor_hour,
    updated_at = NOW();
