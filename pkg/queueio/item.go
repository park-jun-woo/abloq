//ff:type feature=queueio type=schema
//ff:what 큐 항목 1건 — kind/slug/lang/section/priority/payload(문자열 맵), 큐 파일 1개·queue_items 1행과 1:1
package queueio

// Item is one queue entry. Section is a first-class field here even though
// queue_items has no section column — the DB row keeps it inside payload
// (EncodeRows adds it, DecodeRows lifts it back out). Payload values are
// strings only: they round-trip through JSONB byte-identically, which the
// deterministic file serialization (CLI/endpoint diff equality) depends on.
// Timestamps and DB ids never appear here.
type Item struct {
	Kind     string            `json:"kind"`
	Slug     string            `json:"slug"`
	Lang     string            `json:"lang"`
	Section  string            `json:"section"`
	Priority int64             `json:"priority"`
	Payload  map[string]string `json:"payload"`
}
