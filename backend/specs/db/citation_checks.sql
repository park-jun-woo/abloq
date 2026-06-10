-- link rot probe state per citation occurrence. The same URL cited by two
-- articles is two rows — the fix happens per article — so the key is
-- (url, lang, section, slug). status classifies the last probe (ok | hard:
-- 404/410/dead domain | soft: 5xx/timeout); rot confirmation reads
-- consecutive_failures only (>= 3 scans), never the classification, so a
-- transient 404 can recover. status is probe telemetry, not a workflow state
-- machine — no stateDiagram, the XSM-23 deferral comment precedent of
-- queue_items applies. Rows whose citation left the corpus simply stop being
-- updated.
-- @func-managed
CREATE TABLE citation_checks (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    url TEXT NOT NULL,
    lang VARCHAR(35) NOT NULL,
    section TEXT NOT NULL,
    slug TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'ok' CHECK (status IN ('ok', 'hard', 'soft')),
    consecutive_failures BIGINT NOT NULL DEFAULT 0,
    last_checked_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_citation_checks_key ON citation_checks(url, lang, section, slug);
