-- name: SiteFindActiveByName :one
-- The {site} path-param resolver every domain handler runs first: name is
-- the URL-safe slug, only active rows serve. An unregistered or deactivated
-- site reads as no row — the handler's empty-guard turns that into a plain 404.
-- +allow-sensitive
SELECT * FROM sites WHERE name = @name AND active = TRUE;

-- name: SiteListFiltered :many
-- +no-pagination
-- +allow-sensitive
-- Operational/cron lookup (GET /sites). active rides as text ('', 'true',
-- 'false') so one query serves all three states — the cron runners iterate
-- ?active=true. indexnow_key never reaches the response (DDL sensitive tag).
SELECT * FROM sites
WHERE (@active_filter::text = '' OR active = (@active_filter::text)::boolean)
ORDER BY name;

-- name: SiteUpsert :exec
-- Boot-time SSOT sync (cmd/, not a handler): one upsert per declared site,
-- keyed by name. active mirrors the declaration — an explicit
-- "active: false" keeps the row (FK history) but stops it serving;
-- reactivation is simply re-declaring the site active in sites.yaml.
INSERT INTO sites (name, repo_path, queue_export_repo, queue_export_author,
                   queue_export_author_email, gsc_site, gsc_sa_path,
                   cf_log_source, indexnow_key, active)
VALUES (@name, @repo_path, @queue_export_repo, @queue_export_author,
        @queue_export_author_email, @gsc_site, @gsc_sa_path,
        @cf_log_source, @indexnow_key, @active)
ON CONFLICT (name) DO UPDATE SET
    repo_path = EXCLUDED.repo_path,
    queue_export_repo = EXCLUDED.queue_export_repo,
    queue_export_author = EXCLUDED.queue_export_author,
    queue_export_author_email = EXCLUDED.queue_export_author_email,
    gsc_site = EXCLUDED.gsc_site,
    gsc_sa_path = EXCLUDED.gsc_sa_path,
    cf_log_source = EXCLUDED.cf_log_source,
    indexnow_key = EXCLUDED.indexnow_key,
    active = EXCLUDED.active,
    updated_at = NOW();

-- name: SiteDeactivateMissing :exec
-- Boot-time SSOT sync companion: a site that left sites.yaml goes inactive.
-- Rows are never deleted — the FK history (posts, receipts, reports, …)
-- must survive a temporary de-declaration.
UPDATE sites SET active = FALSE, updated_at = NOW()
WHERE active = TRUE AND NOT (name = ANY(@names::text[]));
