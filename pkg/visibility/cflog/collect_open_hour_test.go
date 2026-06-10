//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 안전마진 골든 케이스 — 마진(2h) 안의 열린 시간대 파일은 건드리지 않고 커서도 닫힌 경계까지만 전진
package cflog

import (
	"testing"
	"time"
)

// TestCollectOpenHour pins the late-delivery defense: with now 02:30 and a
// 2h margin the last closed hour is 2026-06-01-23, so the 2026-06-02-00/01
// fixture files stay untouched and the cursor stops at the closed boundary
// (a later collect picks them up).
func TestCollectOpenHour(t *testing.T) {
	root, b := writeRepoFixture(t)
	urls, err := BuildURLMap(root, b)
	if err != nil {
		t.Fatalf("BuildURLMap: %v", err)
	}
	now := time.Date(2026, 6, 2, 2, 30, 0, 0, time.UTC)
	res, err := Collect(DirSource{Root: "testdata/logs"}, urls, nil, now, 2*time.Hour)
	if err != nil {
		t.Fatalf("Collect: %v", err)
	}
	if res.Files != 2 {
		t.Errorf("Files = %d, want 2 (only the 06-01 hours are closed)", res.Files)
	}
	if got := res.Cursors[0].CursorHour; got != "2026-06-01-23" {
		t.Errorf("CursorHour = %q, want 2026-06-01-23", got)
	}
	if res.Total != 8 {
		t.Errorf("Total = %d, want 8 (the 06-02 hits are not ingested yet)", res.Total)
	}
}
