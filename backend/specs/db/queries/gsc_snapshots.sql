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

-- name: GscSnapshotAggPageMonthJson :one
-- Per-page impression/click sums over the report window ('' = the last
-- closed month, UTC). Pages stay full URLs: the page->article attribution
-- happens in Go via the repository URL map (blog.yaml owns the URL rules),
-- so SQL must never LIKE-join them.
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            + interval '1 month - 1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'page', s.page,
           'impressions', s.impressions,
           'clicks', s.clicks
       ) ORDER BY s.page), '[]'::jsonb)::text
FROM (
    SELECT g.page,
           SUM(g.impressions)::BIGINT AS impressions, SUM(g.clicks)::BIGINT AS clicks
    FROM gsc_snapshots g, bounds b
    WHERE g.snap_date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY g.page
) AS s;

-- name: GscSnapshotAggPagePrevMonthJson :one
-- The same aggregate over the previous month's window (trend column).
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            - interval '1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'page', s.page,
           'impressions', s.impressions,
           'clicks', s.clicks
       ) ORDER BY s.page), '[]'::jsonb)::text
FROM (
    SELECT g.page,
           SUM(g.impressions)::BIGINT AS impressions, SUM(g.clicks)::BIGINT AS clicks
    FROM gsc_snapshots g, bounds b
    WHERE g.snap_date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY g.page
) AS s;
