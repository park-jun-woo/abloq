-- Crawl-ingest cursors, one row per log source (cf_logs). cursor_hour is the
-- last UTC hour ("YYYY-MM-DD-HH") ingested wholesale — the cursor is a time
-- boundary, never a file key: CloudFront key suffixes are random, so a
-- key-based start-after cursor would permanently drop late-delivered files.
-- Re-aggregation procedure: resetting a cursor MUST be one operation with
-- deleting the target window (DELETE FROM crawl_hits WHERE hit_date BETWEEN
-- ...) — a bare reset double-accumulates (manual SQL, no API).
CREATE TABLE ingest_cursors (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    site_id BIGINT NOT NULL REFERENCES sites(id),
    source VARCHAR(64) NOT NULL,
    cursor_hour VARCHAR(13) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_ingest_cursors_site_source ON ingest_cursors(site_id, source);
