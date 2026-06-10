//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what lastClosedHour가 마진을 뺀 시각 이전에 끝난 마지막 시간대를 돌려주는지 검증 — 경계 정확성 포함
package cflog

import (
	"testing"
	"time"
)

func TestLastClosedHour(t *testing.T) {
	cases := []struct {
		now    time.Time
		margin time.Duration
		want   string
	}{
		{time.Date(2026, 6, 2, 2, 30, 0, 0, time.UTC), 2 * time.Hour, "2026-06-01-23"},
		{time.Date(2026, 6, 2, 3, 0, 0, 0, time.UTC), 2 * time.Hour, "2026-06-02-00"},
		{time.Date(2026, 6, 2, 2, 59, 59, 0, time.UTC), 2 * time.Hour, "2026-06-01-23"},
		{time.Date(2026, 6, 2, 0, 0, 0, 0, time.UTC), 0, "2026-06-01-23"},
		{time.Date(2026, 6, 2, 12, 0, 0, 0, time.UTC), 24 * time.Hour, "2026-06-01-11"},
	}
	for _, c := range cases {
		if got := lastClosedHour(c.now, c.margin); got != c.want {
			t.Errorf("lastClosedHour(%v, %v) = %q, want %q", c.now, c.margin, got, c.want)
		}
	}
}
