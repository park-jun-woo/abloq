-- status lifecycle: open → exported → consumed → done (DEFAULT + CHECK below).
-- The Mermaid stateDiagram ships in Phase018 with the transition operations —
-- yongol requires every transition label to be an SSaC function (XSM-23);
-- consumed → done belongs to the Phase018 consumers. Until then the
-- transition queries carry WHERE status guards (Phase009).
CREATE TABLE queue_items (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    site_id BIGINT NOT NULL REFERENCES sites(id),
    kind VARCHAR(20) NOT NULL CHECK (kind IN ('refresh', 'evidence', 'cluster')),
    slug TEXT NOT NULL,
    lang VARCHAR(35) NOT NULL,
    payload JSONB NOT NULL DEFAULT '{}', -- @sensitive
    priority BIGINT NOT NULL DEFAULT 0,
    status VARCHAR(20) NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'exported', 'consumed', 'done')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_queue_items_site_status_priority ON queue_items(site_id, status, priority DESC);
