//ff:func feature=queueio type=generator control=sequence
//ff:what EncodeRows가 section·keys를 payload 안에 넣은 적재 JSON을 내고 DecodeRows가 이를 원형 복원하는지 검증 (왕복 동등)
package queueio

import (
	"reflect"
	"testing"
)

func TestEncodeRowsRoundTrip(t *testing.T) {
	items := []Item{{
		Kind: "refresh", Slug: "post-a", Lang: "ko", Section: "tech",
		Priority: 20605, Keys: []string{"ko/tech/post-a", "en/tech/post-a"},
		Payload: map[string]string{"lastmod": "2026-06-05"},
	}}
	rows, err := DecodeRows(EncodeRows(items))
	if err != nil {
		t.Fatalf("DecodeRows: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("want 1 row, got %d", len(rows))
	}
	if !reflect.DeepEqual(rows[0].Item, items[0]) {
		t.Errorf("round trip mismatch:\n got %+v\nwant %+v", rows[0].Item, items[0])
	}
	if string(EncodeRows(nil)) != "[]" {
		t.Errorf("empty encode must be []: %s", EncodeRows(nil))
	}
}
