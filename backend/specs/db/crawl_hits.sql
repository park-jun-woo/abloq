-- @func-managed
CREATE TABLE crawl_hits (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    hit_date DATE NOT NULL,
    bot VARCHAR(64) NOT NULL,
    lang VARCHAR(35) NOT NULL,
    section TEXT NOT NULL,
    slug TEXT NOT NULL,
    hits BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_crawl_hits_key ON crawl_hits(hit_date, bot, lang, section, slug);
