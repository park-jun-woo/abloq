//ff:func feature=visibility type=parser control=iteration dimension=1 topic=report
//ff:what LastClosedYM이 UTC 기준 직전 월을 내는지 검증 — 월초·월말·연 경계 케이스
package report

import (
	"testing"
	"time"
)

func TestLastClosedYM(t *testing.T) {
	cases := []struct {
		now  time.Time
		want string
	}{
		{time.Date(2026, 6, 11, 0, 0, 0, 0, time.UTC), "2026-05"},
		{time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC), "2026-05"},
		{time.Date(2026, 1, 31, 23, 59, 59, 0, time.UTC), "2025-12"},
		{time.Date(2026, 3, 31, 12, 0, 0, 0, time.UTC), "2026-02"},
	}
	for _, tc := range cases {
		if got := LastClosedYM(tc.now); got != tc.want {
			t.Errorf("LastClosedYM(%s): want %s, got %s", tc.now, tc.want, got)
		}
	}
}
