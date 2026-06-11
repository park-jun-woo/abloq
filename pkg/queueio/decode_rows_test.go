//ff:func feature=queueio type=parser control=sequence
//ff:what DecodeRows가 jsonb_agg JSON에서 id를 보존하고 payload의 section·keys를 1급 필드로 승격하는지 검증
package queueio

import (
	"reflect"
	"testing"
)

func TestDecodeRows(t *testing.T) {
	data := `[{"id":3,"kind":"refresh","slug":"post-a","lang":"ko",` +
		`"payload":{"section":"tech","keys":"[\"ko/tech/post-a\",\"en/tech/post-a\"]","lastmod":"2026-06-05"},"priority":20605}]`
	rows, err := DecodeRows([]byte(data))
	if err != nil {
		t.Fatalf("DecodeRows: %v", err)
	}
	r := rows[0]
	if r.ID != 3 || r.Section != "tech" || r.Slug != "post-a" {
		t.Errorf("unexpected row: %+v", r)
	}
	if _, dup := r.Payload["section"]; dup {
		t.Error("section must be lifted out of payload")
	}
	if _, dup := r.Payload["keys"]; dup {
		t.Error("keys must be lifted out of payload")
	}
	if want := []string{"ko/tech/post-a", "en/tech/post-a"}; !reflect.DeepEqual(r.Keys, want) {
		t.Errorf("keys = %v, want %v", r.Keys, want)
	}
	if r.Payload["lastmod"] != "2026-06-05" {
		t.Errorf("payload lost lastmod: %+v", r.Payload)
	}
	legacy := `[{"id":4,"kind":"refresh","slug":"post-b","lang":"ko",` +
		`"payload":{"section":"tech"},"priority":1}]`
	rows, err = DecodeRows([]byte(legacy))
	if err != nil {
		t.Fatalf("DecodeRows legacy: %v", err)
	}
	if rows[0].Keys != nil {
		t.Errorf("legacy row without keys must yield nil: %v", rows[0].Keys)
	}
	if _, err := DecodeRows([]byte("not json")); err == nil {
		t.Error("invalid JSON must error")
	}
}
