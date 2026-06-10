//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 현재 UTC와 안전마진으로 마지막 닫힌 시간대(YYYY-MM-DD-HH) 계산 — 그 시간대의 끝이 now−margin 이전이어야 닫힘
//ff:why CF 로그 배달은 지연될 수 있다(드물게 24h) — 마진 안의 열린 시간대를 건드리면 늦게 배달된 파일이 영구 누락된다. 마진은 env로 조정 가능, 기본 2h (Phase012)
package cflog

import "time"

// lastClosedHour returns the latest UTC hour H whose end (H+1h) is at or
// before now-margin: every log file of H is assumed delivered by then.
func lastClosedHour(now time.Time, margin time.Duration) string {
	end := now.UTC().Add(-margin)
	return end.Add(-time.Hour).Truncate(time.Hour).Format("2006-01-02-15")
}
