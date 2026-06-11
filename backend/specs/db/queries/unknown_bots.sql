-- name: UnknownBotUpsertFromJson :exec
-- Accumulate unknown-bot candidates per unique (site, UA): hits add up, the
-- first/last sighting bounds widen monotonically. Site-scoped on purpose —
-- UA knowledge is global, but "observed in this site's logs" is the report's
-- fact unit.
WITH incoming AS (
    SELECT (e->>'ua')::TEXT                  AS ua,
           (e->>'hits')::BIGINT              AS hits,
           (e->>'first_seen')::TIMESTAMPTZ   AS first_seen,
           (e->>'last_seen')::TIMESTAMPTZ    AS last_seen
    FROM jsonb_array_elements(@bots_json::jsonb) AS e
)
INSERT INTO unknown_bots (site_id, ua, hits, first_seen, last_seen)
SELECT @site_id, i.ua, i.hits, i.first_seen, i.last_seen
FROM incoming i
ON CONFLICT (site_id, ua) DO UPDATE
SET hits = unknown_bots.hits + EXCLUDED.hits,
    first_seen = LEAST(unknown_bots.first_seen, EXCLUDED.first_seen),
    last_seen = GREATEST(unknown_bots.last_seen, EXCLUDED.last_seen);

-- name: UnknownBotAggUasJson :one
-- Unknown-bot candidates for the report (dictionary-update input): every
-- accumulated UA with its hit count, busiest first. Not windowed — the list
-- is a standing to-do, not a monthly metric.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'ua', ua,
           'hits', hits
       ) ORDER BY hits DESC, ua), '[]'::jsonb)::text
FROM unknown_bots
WHERE site_id = @site_id;
