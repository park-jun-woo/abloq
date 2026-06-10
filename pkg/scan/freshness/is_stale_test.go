//ff:func feature=scan type=rule control=iteration dimension=1
//ff:what isStale의 경계(정확히 threshold 경과는 비초과)·초과·미파싱 케이스 검증
package freshness

import (
	"testing"
	"time"
)

func TestIsStale(t *testing.T) {
	now := time.Date(2026, 6, 11, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		name    string
		lastmod string
		days    int
		want    bool
	}{
		{"stale", "2026-06-05", 1, true},
		{"exactly at threshold", "2026-06-10", 1, false},
		{"fresh", "2026-06-11", 1, false},
		{"unparseable", "n/a", 1, false},
		{"empty", "", 1, false},
	}
	for _, tc := range cases {
		if got := isStale(tc.lastmod, tc.days, now); got != tc.want {
			t.Errorf("%s: want %v, got %v", tc.name, tc.want, got)
		}
	}
}
