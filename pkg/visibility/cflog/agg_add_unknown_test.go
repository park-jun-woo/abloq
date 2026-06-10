//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what addUnknown이 휴리스틱 통과 UA만 누적하고 최초/최종 목격 시각을 양방향 갱신하는지 검증
package cflog

import (
	"testing"
	"time"
)

func TestAggAddUnknown(t *testing.T) {
	agg := NewAgg(nil)
	t1 := time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC)
	t0 := t1.Add(-time.Hour)
	t2 := t1.Add(time.Hour)
	ua := "PetalBot/1.0"
	agg.addUnknown(Record{When: t1, UA: ua})
	agg.addUnknown(Record{When: t2, UA: ua})
	agg.addUnknown(Record{When: t0, UA: ua})
	agg.addUnknown(Record{When: t1, UA: "Mozilla/5.0 Chrome/125 Safari/537.36"})
	agg.addUnknown(Record{When: t1, UA: "-"})
	rows := agg.UnknownRows()
	if len(rows) != 1 {
		t.Fatalf("rows = %+v, want only the PetalBot UA", rows)
	}
	if rows[0].Hits != 3 || rows[0].FirstSeen != "2026-06-01T09:00:00Z" || rows[0].LastSeen != "2026-06-01T11:00:00Z" {
		t.Errorf("row = %+v", rows[0])
	}
}
