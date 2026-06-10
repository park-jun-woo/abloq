//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 누적된 미지 봇 후보를 UA 사전순 행 목록으로 — 시각은 UTC RFC3339, 결정적 출력
package cflog

import (
	"sort"
	"time"
)

// UnknownRows flattens the unknown-bot accumulation into rows sorted by UA.
func (a *Agg) UnknownRows() []UnknownRow {
	rows := make([]UnknownRow, 0, len(a.unknown))
	for ua, u := range a.unknown {
		rows = append(rows, UnknownRow{
			UA:        ua,
			Hits:      u.Hits,
			FirstSeen: u.First.UTC().Format(time.RFC3339),
			LastSeen:  u.Last.UTC().Format(time.RFC3339),
		})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].UA < rows[j].UA })
	return rows
}
