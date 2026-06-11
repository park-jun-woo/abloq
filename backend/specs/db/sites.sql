-- The FK anchor of the multisite instance (Phase020): one row per declared
-- blog site. The SSOT is the deploy-side sites.yaml (SITES_YAML_PATH) — boot
-- upserts rows by name and deactivates rows that left the SSOT (never
-- deletes: the FK history must survive). name is the URL-safe {site} path
-- parameter. File paths are not secrets (the mounted file content is);
-- indexnow_key is the one credential value and stays out of API responses.
-- Without SITES_YAML_PATH, a BLOG_REPO_PATH-only boot synthesizes one
-- name='default' row from the legacy per-site env 8종 (backward compat).
CREATE TABLE sites (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    repo_path TEXT NOT NULL,
    queue_export_repo TEXT NOT NULL DEFAULT '',
    queue_export_author TEXT NOT NULL DEFAULT '',
    queue_export_author_email TEXT NOT NULL DEFAULT '',
    gsc_site TEXT NOT NULL DEFAULT '',
    gsc_sa_path TEXT NOT NULL DEFAULT '',
    cf_log_source TEXT NOT NULL DEFAULT '',
    indexnow_key TEXT NOT NULL DEFAULT '', -- @sensitive
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_sites_name ON sites(name);
