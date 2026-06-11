//ff:type feature=queueio type=schema
//ff:what 큐 항목 1건 — kind/slug/lang/section/priority/keys(전 언어 키)/payload(문자열 맵), 큐 파일 1개·queue_items 1행과 1:1
package queueio

// Item is one queue entry. Section is a first-class field here even though
// queue_items has no section column — the DB row keeps it inside payload
// (EncodeRows adds it, DecodeRows lifts it back out). Keys carries the
// gate-contract join key for every declared language (Phase018: consumer
// quests resync translations, so honest-lastmod must allow every language's
// lastmod update from one queue file); it rides inside payload the same way
// section does. Payload values are strings only: they round-trip through
// JSONB byte-identically, which the deterministic file serialization
// (CLI/endpoint diff equality) depends on. Timestamps and DB ids never
// appear here.
type Item struct {
	Kind     string            `json:"kind"`
	Slug     string            `json:"slug"`
	Lang     string            `json:"lang"`
	Section  string            `json:"section"`
	Priority int64             `json:"priority"`
	Keys     []string          `json:"keys,omitempty"`
	Payload  map[string]string `json:"payload"`
}
