-- abloqd v0.1.x(단일 사이트) → v0.2.0(멀티사이트) 기존 DB 마이그레이션 (Phase020)
--
-- 대상: 이미 운용 중인 단일 사이트 배포의 postgres. 새 설치는 이 파일이
-- 필요 없다 — backend/arts/db/migrations/0001_initial.up.sql 이 신규 기준
-- 스키마(멀티사이트)를 그대로 만든다.
--
-- 절차: ① abloqd 중지 → ② psql -f migrate-0.2.0-multisite.sql → ③ 새
-- 이미지로 기동. 기동 동기화가 default 행의 값(repo_path·키들)을
-- env/sites.yaml 선언으로 덮어쓰므로 여기서는 행 존재만 보장하면 된다.
--
-- 규약: 기존 데이터는 전부 name='default' 사이트로 백필한다 — id 하드코딩
-- 금지, 반드시 (SELECT id FROM sites WHERE name='default') 서브쿼리로.
BEGIN;

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
    indexnow_key TEXT NOT NULL DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_sites_name ON sites(name);

-- 기존 데이터의 귀속처. repo_path '/blog'은 compose 컨벤션 임시값 —
-- 기동 동기화(BLOG_REPO_PATH 또는 sites.yaml)가 실값으로 덮어쓴다.
INSERT INTO sites (name, repo_path)
VALUES ('default', '/blog')
ON CONFLICT DO NOTHING;

-- posts: 유니크 (lang, section, slug) → (site_id, lang, section, slug)
ALTER TABLE posts ADD COLUMN site_id BIGINT;
UPDATE posts SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE posts ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE posts ADD CONSTRAINT posts_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_posts_lang_section_slug;
CREATE UNIQUE INDEX idx_posts_site_lang_section_slug ON posts(site_id, lang, section, slug);

-- receipts: 멱등키 (deploy_id, kind, target) → (site_id, deploy_id, kind, target)
ALTER TABLE receipts ADD COLUMN site_id BIGINT;
UPDATE receipts SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE receipts ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE receipts ADD CONSTRAINT receipts_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_receipts_deploy_kind_target;
DROP INDEX idx_receipts_kind_target;
CREATE UNIQUE INDEX idx_receipts_site_deploy_kind_target ON receipts(site_id, deploy_id, kind, target);
CREATE INDEX idx_receipts_site_kind_target ON receipts(site_id, kind, target);

-- queue_items: 유니크 없음 — site_id + 조회 인덱스만
ALTER TABLE queue_items ADD COLUMN site_id BIGINT;
UPDATE queue_items SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE queue_items ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE queue_items ADD CONSTRAINT queue_items_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_queue_items_status_priority;
CREATE INDEX idx_queue_items_site_status_priority ON queue_items(site_id, status, priority DESC);

-- crawl_hits: 유니크 (hit_date, bot, lang, section, slug) → 선두 site_id
ALTER TABLE crawl_hits ADD COLUMN site_id BIGINT;
UPDATE crawl_hits SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE crawl_hits ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE crawl_hits ADD CONSTRAINT crawl_hits_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_crawl_hits_key;
CREATE UNIQUE INDEX idx_crawl_hits_key ON crawl_hits(site_id, hit_date, bot, lang, section, slug);

-- unknown_bots: 유니크 (ua) → (site_id, ua)
ALTER TABLE unknown_bots ADD COLUMN site_id BIGINT;
UPDATE unknown_bots SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE unknown_bots ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE unknown_bots ADD CONSTRAINT unknown_bots_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_unknown_bots_ua;
CREATE UNIQUE INDEX idx_unknown_bots_site_ua ON unknown_bots(site_id, ua);

-- ingest_cursors: 유니크 (source) → (site_id, source)
ALTER TABLE ingest_cursors ADD COLUMN site_id BIGINT;
UPDATE ingest_cursors SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE ingest_cursors ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE ingest_cursors ADD CONSTRAINT ingest_cursors_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_ingest_cursors_source;
CREATE UNIQUE INDEX idx_ingest_cursors_site_source ON ingest_cursors(site_id, source);

-- gsc_snapshots: 유니크 (snap_date, page) → 선두 site_id
ALTER TABLE gsc_snapshots ADD COLUMN site_id BIGINT;
UPDATE gsc_snapshots SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE gsc_snapshots ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE gsc_snapshots ADD CONSTRAINT gsc_snapshots_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_gsc_snapshots_key;
CREATE UNIQUE INDEX idx_gsc_snapshots_key ON gsc_snapshots(site_id, snap_date, page);

-- citation_queries: 유니크 없음 — site_id + 조회 인덱스만.
-- citation_samples는 citation_queries FK 경유로 사이트가 결정된다(무변경).
ALTER TABLE citation_queries ADD COLUMN site_id BIGINT;
UPDATE citation_queries SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE citation_queries ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE citation_queries ADD CONSTRAINT citation_queries_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_citation_queries_post;
CREATE INDEX idx_citation_queries_site_post ON citation_queries(site_id, lang, section, slug);

-- citation_checks: 유니크 (url, lang, section, slug) → 선두 site_id
ALTER TABLE citation_checks ADD COLUMN site_id BIGINT;
UPDATE citation_checks SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE citation_checks ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE citation_checks ADD CONSTRAINT citation_checks_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_citation_checks_key;
CREATE UNIQUE INDEX idx_citation_checks_key ON citation_checks(site_id, url, lang, section, slug);

-- reports: 유니크 (ym) → (site_id, ym) — 월간 리포트가 사이트당 1개
ALTER TABLE reports ADD COLUMN site_id BIGINT;
UPDATE reports SET site_id = (SELECT id FROM sites WHERE name = 'default');
ALTER TABLE reports ALTER COLUMN site_id SET NOT NULL;
ALTER TABLE reports ADD CONSTRAINT reports_site_id_fkey FOREIGN KEY (site_id) REFERENCES sites(id);
DROP INDEX idx_reports_ym;
CREATE UNIQUE INDEX idx_reports_site_ym ON reports(site_id, ym);

-- users · refresh_tokens: 글로벌 유지 — 무변경.

COMMIT;
