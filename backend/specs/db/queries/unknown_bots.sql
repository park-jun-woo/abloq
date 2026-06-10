-- name: UnknownBotUpsertFromJson :exec
-- Accumulate unknown-bot candidates per unique UA: hits add up, the
-- first/last sighting bounds widen monotonically.
WITH incoming AS (
    SELECT (e->>'ua')::TEXT                  AS ua,
           (e->>'hits')::BIGINT              AS hits,
           (e->>'first_seen')::TIMESTAMPTZ   AS first_seen,
           (e->>'last_seen')::TIMESTAMPTZ    AS last_seen
    FROM jsonb_array_elements(@bots_json::jsonb) AS e
)
INSERT INTO unknown_bots (ua, hits, first_seen, last_seen)
SELECT i.ua, i.hits, i.first_seen, i.last_seen
FROM incoming i
ON CONFLICT (ua) DO UPDATE
SET hits = unknown_bots.hits + EXCLUDED.hits,
    first_seen = LEAST(unknown_bots.first_seen, EXCLUDED.first_seen),
    last_seen = GREATEST(unknown_bots.last_seen, EXCLUDED.last_seen);
