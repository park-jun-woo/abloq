-- @func-managed
CREATE TABLE citation_samples (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    citation_queries_id BIGINT NOT NULL REFERENCES citation_queries(id),
    engine VARCHAR(40) NOT NULL,
    cited BOOLEAN NOT NULL DEFAULT FALSE,
    evidence TEXT NOT NULL DEFAULT '',
    extractor_version VARCHAR(40) NOT NULL DEFAULT '', -- @backfill default=''
    run_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_citation_samples_query ON citation_samples(citation_queries_id, run_at);
