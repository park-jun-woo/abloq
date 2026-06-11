-- name: QueueItemInsertMissingFromJson :exec
-- Batch insert of scan candidates. The dedup guard compares
-- payload->>'section' because queue_items has no section column and the
-- posts unique key is (site_id, lang, section, slug): dropping section would
-- silently skip one of two same-slug articles in different sections.
-- consumed|done rows do not block — a consumed article may legitimately go
-- stale again. The guard is site-scoped: the same article key on another
-- site is a different work item.
WITH incoming AS (
    SELECT (e->>'kind')::TEXT              AS kind,
           (e->>'slug')::TEXT              AS slug,
           (e->>'lang')::TEXT              AS lang,
           COALESCE(e->'payload', '{}'::jsonb) AS payload,
           (e->>'priority')::BIGINT        AS priority
    FROM jsonb_array_elements(@items_json::jsonb) AS e
)
INSERT INTO queue_items (site_id, kind, slug, lang, payload, priority)
SELECT @site_id, i.kind, i.slug, i.lang, i.payload, i.priority
FROM incoming i
WHERE NOT EXISTS (
    SELECT 1 FROM queue_items q
    WHERE q.site_id = @site_id
      AND q.kind = i.kind
      AND q.slug = i.slug
      AND q.lang = i.lang
      AND q.payload->>'section' = i.payload->>'section'
      AND q.status IN ('open', 'exported')
);

-- name: QueueItemAggOpenJson :one
-- Open items as a JSON text scalar for the exporter func (the func never
-- sees db model types). Priority DESC mirrors the export ordering; empty set
-- must be '[]', not NULL.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'id', id,
           'kind', kind,
           'slug', slug,
           'lang', lang,
           'payload', payload,
           'priority', priority
       ) ORDER BY priority DESC, id), '[]'::jsonb)::text
FROM queue_items
WHERE site_id = @site_id AND status = 'open';

-- name: QueueItemAggExportedJson :one
-- Exported items for the consumed sync: the exporter recomputes each row's
-- file name (forward only) and checks existence in the fresh work clone.
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'id', id,
           'kind', kind,
           'slug', slug,
           'lang', lang,
           'payload', payload,
           'priority', priority
       ) ORDER BY priority DESC, id), '[]'::jsonb)::text
FROM queue_items
WHERE site_id = @site_id AND status = 'exported';

-- name: QueueItemMarkExportedFromJson :exec
-- open → exported after a successful push. The WHERE status guard preserves
-- the state-machine semantics until the Phase018 stateDiagram declares the
-- transitions (XSM-23 needs every label to be an SSaC function).
UPDATE queue_items SET status = 'exported', updated_at = NOW()
WHERE site_id = @site_id
  AND status = 'open'
  AND id IN (SELECT value::bigint FROM jsonb_array_elements_text(@id_list_json::jsonb));

-- name: QueueItemMarkConsumedFromJson :exec
-- exported → consumed when the agent's consumption commit deleted the file.
UPDATE queue_items SET status = 'consumed', updated_at = NOW()
WHERE site_id = @site_id
  AND status = 'exported'
  AND id IN (SELECT value::bigint FROM jsonb_array_elements_text(@id_list_json::jsonb));

-- name: QueueItemListFiltered :many
-- +no-pagination
-- +allow-sensitive
-- Operational lookup; empty-string filters mean "no filter". payload stays
-- out of the JSON response via the DDL sensitive tag.
SELECT * FROM queue_items
WHERE site_id = @site_id
  AND (@kind::text = '' OR kind = @kind::text)
  AND (@status::text = '' OR status = @status::text)
ORDER BY priority DESC, id;

-- name: QueueItemAggMonthCountsJson :one
-- Queue intake summary of the report window: rows created (UTC date) inside
-- the 30 days ending on the last day of @ym, counted per (kind, status).
-- created_at is DB NOW(), so an explicit past ym deterministically reads 0
-- — the Hurl oracle; non-zero goldens live in the pkg tests.
WITH bounds AS (
    SELECT (date_trunc('month', CASE WHEN @ym::text = ''
                THEN date_trunc('month', now() AT TIME ZONE 'utc') - interval '1 month'
                ELSE to_date(@ym::text || '-01', 'YYYY-MM-DD')::timestamp END)
            + interval '1 month - 1 day')::date AS month_end
)
SELECT COALESCE(jsonb_agg(jsonb_build_object(
           'kind', s.kind,
           'status', s.status,
           'count', s.cnt
       ) ORDER BY s.kind, s.status), '[]'::jsonb)::text
FROM (
    SELECT qi.kind, qi.status, COUNT(*)::BIGINT AS cnt
    FROM queue_items qi, bounds b
    WHERE qi.site_id = @site_id
      AND (qi.created_at AT TIME ZONE 'utc')::date BETWEEN b.month_end - 29 AND b.month_end
    GROUP BY qi.kind, qi.status
) AS s;
