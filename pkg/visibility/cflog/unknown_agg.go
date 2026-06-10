//ff:type feature=visibility type=schema topic=crawl
//ff:what 미지 봇 UA 1개의 누적 상태 — 히트 수와 최초/최종 목격 시각
package cflog

import "time"

// unknownAgg accumulates one unknown UA's sightings.
type unknownAgg struct {
	Hits  int64
	First time.Time
	Last  time.Time
}
