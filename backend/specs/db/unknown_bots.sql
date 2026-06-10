-- Unknown-bot candidates: UAs the crawl ingest saw that are not in the
-- pkg/bots dictionary but match the bot heuristic (bot/crawler/spider tokens
-- or non-browser client patterns). Dictionary-update input for the operator;
-- ordinary browser UAs and empty UAs never land here.
CREATE TABLE unknown_bots (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ua TEXT NOT NULL,
    hits BIGINT NOT NULL DEFAULT 0,
    first_seen TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_seen TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_unknown_bots_ua ON unknown_bots(ua);
