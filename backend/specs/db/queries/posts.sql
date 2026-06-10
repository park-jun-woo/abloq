-- name: PostSyncFromJson :exec
WITH incoming AS (
    -- jsonb_array_elements + per-field extraction instead of
    -- jsonb_to_recordset: sqlc cannot resolve a recordset coldeflist
    -- (it collapses `SELECT *` into one record column and rejects
    -- qualified references), while plain jsonb operators analyze fine.
    SELECT (e->>'lang')::TEXT            AS lang,
           (e->>'section')::TEXT         AS section,
           (e->>'slug')::TEXT            AS slug,
           (e->>'title')::TEXT           AS title,
           (e->>'date')::TEXT            AS "date",
           (e->>'lastmod')::TEXT         AS lastmod,
           (e->>'word_count')::BIGINT    AS word_count,
           COALESCE(e->'tags', '[]'::jsonb) AS tags,
           (e->>'internal_links')::BIGINT AS internal_links,
           (e->>'source_count')::BIGINT  AS source_count,
           (e->>'url')::TEXT             AS url
    FROM jsonb_array_elements(@entries_json::jsonb) AS e
), removed AS (
    DELETE FROM posts
    WHERE (lang, section, slug) NOT IN (SELECT lang, section, slug FROM incoming)
)
INSERT INTO posts (lang, section, slug, title, "date", lastmod, word_count,
                   tags, internal_links, source_count, url)
SELECT lang, section, slug, title, "date", lastmod, word_count,
       tags, internal_links, source_count, url
FROM incoming
ON CONFLICT (lang, section, slug) DO UPDATE SET
    title = EXCLUDED.title,
    "date" = EXCLUDED."date",
    lastmod = EXCLUDED.lastmod,
    word_count = EXCLUDED.word_count,
    tags = EXCLUDED.tags,
    internal_links = EXCLUDED.internal_links,
    source_count = EXCLUDED.source_count,
    url = EXCLUDED.url,
    updated_at = NOW();

-- name: PostListAll :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM posts ORDER BY lang, section, slug;
