-- name: ReportUpsert :exec
-- One report row per (site, ym): regeneration replaces (the report derives
-- from ym-anchored aggregates, so a re-run with the same data is
-- byte-identical and the update converges).
INSERT INTO reports (site_id, ym, markdown, report_json)
VALUES (@site_id, @ym::text, @markdown::text, @report_json::text)
ON CONFLICT (site_id, ym) DO UPDATE
SET markdown = EXCLUDED.markdown,
    report_json = EXCLUDED.report_json,
    updated_at = NOW();

-- name: ReportFindByYm :one
-- Lookup truth of GET /sites/{site}/reports/monthly/{ym}.
SELECT * FROM reports WHERE site_id = @site_id AND ym = @ym::text;
