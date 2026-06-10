-- name: ReportUpsert :exec
-- One report row per ym: regeneration replaces (the report derives from
-- ym-anchored aggregates, so a re-run with the same data is byte-identical
-- and the update converges).
INSERT INTO reports (ym, markdown, report_json)
VALUES (@ym::text, @markdown::text, @report_json::text)
ON CONFLICT (ym) DO UPDATE
SET markdown = EXCLUDED.markdown,
    report_json = EXCLUDED.report_json,
    updated_at = NOW();

-- name: ReportFindByYm :one
-- Lookup truth of GET /reports/monthly/{ym}.
SELECT * FROM reports WHERE ym = @ym::text;
