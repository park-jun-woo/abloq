-- name: IngestCursorAggAllJson :one
-- Cursor state as a JSON text scalar for the ingest func (the func never
-- sees db model types — ScanEvidence precedent). Empty set must be '[]'.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'source', source,
           'cursor_hour', cursor_hour
       ) ORDER BY source), '[]'::jsonb)::text
FROM ingest_cursors;

-- name: IngestCursorUpsertFromJson :exec
-- Advance cursors to the closed-hour boundaries the ingest just covered.
WITH incoming AS (
    SELECT (e->>'source')::TEXT      AS source,
           (e->>'cursor_hour')::TEXT AS cursor_hour
    FROM jsonb_array_elements(@cursors_json::jsonb) AS e
)
INSERT INTO ingest_cursors (source, cursor_hour)
SELECT i.source, i.cursor_hour
FROM incoming i
ON CONFLICT (source) DO UPDATE
SET cursor_hour = EXCLUDED.cursor_hour,
    updated_at = NOW();
