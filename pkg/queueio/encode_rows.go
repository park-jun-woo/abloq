//ff:func feature=queueio type=generator control=iteration dimension=1
//ff:what 후보 항목 → queue_items 배치 적재 JSON — section을 payload 안으로 (queue_items에 section 컬럼이 없음)
package queueio

import "encoding/json"

// EncodeRows turns scan candidates into the JSON array that
// QueueItem.InsertMissingFromJson consumes. Section moves inside payload:
// the dedup query compares payload->>'section' because the posts unique key
// is (lang, section, slug) — dropping section would silently skip one of two
// same-slug articles in different sections.
func EncodeRows(items []Item) []byte {
	rows := make([]map[string]any, 0, len(items))
	for _, it := range items {
		payload := map[string]string{"section": it.Section}
		for k, v := range it.Payload {
			payload[k] = v
		}
		rows = append(rows, map[string]any{
			"kind":     it.Kind,
			"slug":     it.Slug,
			"lang":     it.Lang,
			"payload":  payload,
			"priority": it.Priority,
		})
	}
	// Marshal cannot fail: the structure is maps of strings and int64 only.
	data, _ := json.Marshal(rows)
	return data
}
