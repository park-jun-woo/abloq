//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what TotalHits가 hits+md_hits 총합을 내는지 검증
package cflog

import "testing"

func TestTotalHits(t *testing.T) {
	rows := []HitRow{{Hits: 2, MDHits: 1}, {Hits: 0, MDHits: 3}}
	if got := TotalHits(rows); got != 6 {
		t.Errorf("totalHits = %d, want 6", got)
	}
	if got := TotalHits(nil); got != 0 {
		t.Errorf("TotalHits(nil) = %d, want 0", got)
	}
}
