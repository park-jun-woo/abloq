-- Monthly visibility reports, one row per ym (YYYY-MM). The DB row is the
-- lookup truth GET /reports/monthly/{ym} reads; the git commit of the
-- markdown (reports/<ym>.md in the blog repository) is a publication copy.
-- report_json is TEXT (not JSONB): the machine output is stored verbatim,
-- byte-identical to what pkg/visibility/report rendered.
CREATE TABLE reports (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    site_id BIGINT NOT NULL REFERENCES sites(id),
    ym VARCHAR(7) NOT NULL,
    markdown TEXT NOT NULL,
    report_json TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_reports_site_ym ON reports(site_id, ym);
