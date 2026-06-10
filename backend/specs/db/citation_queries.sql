-- @func-managed
CREATE TABLE citation_queries (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    lang VARCHAR(35) NOT NULL,
    section TEXT NOT NULL,
    slug TEXT NOT NULL,
    query_text TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_citation_queries_post ON citation_queries(lang, section, slug);
