# Abloqd

abloq operations backend — content index, receipts, queues, visibility metrics.
**v0.2.0 — multisite**: one instance serves every site declared in the
deploy-side `sites.yaml` (`SITES_YAML_PATH`); every domain endpoint lives
under `/sites/{site}/…` and `GET /sites` lists the registry.

## Development

`yongol generate specs arts` to generate code artifacts.

Run `go build ./...` inside `arts/backend/`, then start the server.

Boot-time site registry sync:

- `SITES_YAML_PATH` set → strict-parse + validate `sites.yaml`
  (`pkg/sitesyaml`), upsert every declared site by name, deactivate rows
  that left the SSOT (never delete — FK history). Any diagnostic refuses
  the boot.
- else `BLOG_REPO_PATH` set → synthesize the single `default` site from the
  legacy per-site env 8종 (`BLOG_REPO_PATH`, `QUEUE_EXPORT_REPO_URL`
  `·AUTHOR·AUTHOR_EMAIL`, `GSC_SITE_URL`, `GSC_SA_JSON_PATH`,
  `CF_LOG_SOURCE`, `INDEXNOW_KEY`) — single-site backward compat; API
  paths still move to `/sites/default/…`.

## Tests

- `backend/scripts/hurl-test/run.sh` — full Hurl suite on a throwaway
  postgres: the shared instance boots from a 2-site (+1 inactive)
  sites.yaml fixture and `scenario-multisite.hurl` pins the isolation
  oracles; the cluster instance pins the BLOG_REPO_PATH-only boot.
- `backend/scripts/compose-cron-smoke/run.sh` — cron shell-expansion smoke
  (every runner iterates `GET /sites?active_filter=true` and fans out per
  site).
- `backend/scripts/rehearsal/run.sh` — one full operating-loop revolution
  on fixtures (single `default` site).

Operations manual: `docs/operations.md` (§5 multisite).
