CREATE TABLE posts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    lang VARCHAR(35) NOT NULL,
    section TEXT NOT NULL,
    slug TEXT NOT NULL,
    title TEXT NOT NULL,
    date TEXT NOT NULL DEFAULT '',
    lastmod TEXT NOT NULL DEFAULT '',
    word_count BIGINT NOT NULL DEFAULT 0,
    tags JSONB NOT NULL DEFAULT '[]', -- @sensitive
    internal_links BIGINT NOT NULL DEFAULT 0,
    source_count BIGINT NOT NULL DEFAULT 0,
    url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_posts_lang_section_slug ON posts(lang, section, slug);
