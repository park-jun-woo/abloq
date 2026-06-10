-- status lifecycle: pending → done | failed | deferred (archiver receipt-table
-- queue — Phase008). pending rows are the work queue; POST /archive/process
-- executes them, POST /receipts/retry rearms failed/deferred back to pending.
CREATE TABLE receipts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    deploy_id TEXT NOT NULL DEFAULT '',
    kind VARCHAR(20) NOT NULL CHECK (kind IN ('wayback', 'indexnow', 'gsc_index')),
    target TEXT NOT NULL,
    request JSONB NOT NULL DEFAULT '{}', -- @sensitive
    response JSONB NOT NULL DEFAULT '{}', -- @sensitive
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'done', 'failed', 'deferred')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Idempotency key: one receipt per (deploy, kind, target). The same URL gets a
-- fresh receipt on every later deploy that changes it — evidence per change.
CREATE UNIQUE INDEX idx_receipts_deploy_kind_target ON receipts(deploy_id, kind, target);
CREATE INDEX idx_receipts_kind_target ON receipts(kind, target);
