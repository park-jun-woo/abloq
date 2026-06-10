//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what UnknownRows가 UA 사전순으로 행을 고정하고 RFC3339 UTC 시각을 내는지 검증
package cflog

import (
	"testing"
	"time"
)

func TestUnknownRows(t *testing.T) {
	agg := NewAgg(nil)
	when := time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC)
	agg.addUnknown(Record{When: when, UA: "zspider/1.0"})
	agg.addUnknown(Record{When: when, UA: "curl/8.5.0"})
	rows := agg.UnknownRows()
	if len(rows) != 2 || rows[0].UA != "curl/8.5.0" || rows[1].UA != "zspider/1.0" {
		t.Fatalf("rows = %+v, want UA-sorted pair", rows)
	}
	if rows[0].FirstSeen != "2026-06-01T12:00:00Z" {
		t.Errorf("FirstSeen = %q", rows[0].FirstSeen)
	}
}
