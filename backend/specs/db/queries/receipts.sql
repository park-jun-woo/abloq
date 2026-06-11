-- name: ReceiptAggByDeployJson :one
-- Existing (kind, target) pairs of one deploy as a JSON text scalar — the
-- webhook func filters already-receipted pairs out (idempotent re-webhook).
-- Empty set must be '[]', not NULL: COALESCE is mandatory.
SELECT COALESCE(jsonb_agg(jsonb_build_object('kind', kind, 'target', target)
                          ORDER BY id), '[]'::jsonb)::text
FROM receipts
WHERE site_id = @site_id AND deploy_id = @deploy_id;

-- name: ReceiptAggPendingJson :one
-- Pending receipts as a JSON text scalar for the processor func. posts is
-- joined on the canonical URL (same site) to supply date/lastmod — the GSC
-- quota split prioritises new posts (date == lastmod) over updated ones
-- inside pkg/archive.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'deploy_id', r.deploy_id,
           'kind', r.kind,
           'target', r.target,
           'date', COALESCE(p.date, ''),
           'lastmod', COALESCE(p.lastmod, '')
       ) ORDER BY r.id), '[]'::jsonb)::text
FROM receipts r
LEFT JOIN posts p ON p.site_id = r.site_id AND p.url = r.target
WHERE r.site_id = @site_id AND r.status = 'pending';

-- name: ReceiptUpsertFromJson :exec
-- Batch upsert on the idempotency key (site_id, deploy_id, kind, target).
-- ON CONFLICT must be DO UPDATE: the processor records the final status over
-- the pending row — DO NOTHING would leave pending rows pending forever.
WITH incoming AS (
    SELECT (e->>'deploy_id')::TEXT            AS deploy_id,
           (e->>'kind')::TEXT                 AS kind,
           (e->>'target')::TEXT               AS target,
           COALESCE(e->'request', '{}'::jsonb)  AS request,
           COALESCE(e->'response', '{}'::jsonb) AS response,
           (e->>'status')::TEXT               AS status
    FROM jsonb_array_elements(@items_json::jsonb) AS e
)
INSERT INTO receipts (site_id, deploy_id, kind, target, request, response, status)
SELECT @site_id, deploy_id, kind, target, request, response, status
FROM incoming
ON CONFLICT (site_id, deploy_id, kind, target) DO UPDATE SET
    request = EXCLUDED.request,
    response = EXCLUDED.response,
    status = EXCLUDED.status;

-- name: ReceiptRearmFromFilter :exec
-- failed/deferred → pending (the next /archive/process executes them).
-- Empty-string filters mean "no filter".
UPDATE receipts SET status = 'pending'
WHERE site_id = @site_id
  AND status IN ('failed', 'deferred')
  AND (@deploy_id::text = '' OR deploy_id = @deploy_id::text)
  AND (@kind::text = '' OR kind = @kind::text);

-- name: ReceiptCountPending :one
SELECT COUNT(*) FROM receipts WHERE site_id = @site_id AND status = 'pending';

-- name: ReceiptListFiltered :many
-- +no-pagination
-- +allow-sensitive
-- Gate-facing receipt lookup; empty-string filters mean "no filter".
-- request/response stay out of the JSON response via the DDL sensitive tags.
SELECT * FROM receipts
WHERE site_id = @site_id
  AND (@target::text = '' OR target = @target::text)
  AND (@kind::text = '' OR kind = @kind::text)
  AND (@deploy_id::text = '' OR deploy_id = @deploy_id::text)
  AND (@status::text = '' OR status = @status::text)
ORDER BY id;
