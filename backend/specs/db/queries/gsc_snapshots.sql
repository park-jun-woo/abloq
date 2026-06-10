-- name: GscSnapshotMaxDateText :one
-- The GSC ingest cursor, derived from the snapshots themselves — no cursor
-- table. Empty table must be '', not NULL (receipts COALESCE precedent):
-- the func treats '' as "first collection, use the lookback window".
SELECT COALESCE(MAX(snap_date)::text, '')::text FROM gsc_snapshots;

-- name: GscSnapshotUpsertFromJson :exec
-- Batch upsert of one collection's (day, page) rows. GSC reports final
-- numbers for closed days, so ON CONFLICT replaces (never adds) — a
-- re-collection of the same day converges instead of accumulating.
WITH incoming AS (
    SELECT (e->>'snap_date')::DATE                 AS snap_date,
           (e->>'page')::TEXT                      AS page,
           (e->>'impressions')::BIGINT             AS impressions,
           (e->>'clicks')::BIGINT                  AS clicks,
           (e->>'avg_position')::DOUBLE PRECISION  AS avg_position
    FROM jsonb_array_elements(@rows_json::jsonb) AS e
)
INSERT INTO gsc_snapshots (snap_date, page, impressions, clicks, avg_position)
SELECT i.snap_date, i.page, i.impressions, i.clicks, i.avg_position
FROM incoming i
ON CONFLICT (snap_date, page) DO UPDATE
SET impressions = EXCLUDED.impressions,
    clicks = EXCLUDED.clicks,
    avg_position = EXCLUDED.avg_position;
