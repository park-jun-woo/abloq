//ff:func feature=scan type=parser control=sequence topic=cluster
//ff:what date 스칼라 → epoch 일수 — RFC3339 또는 YYYY-MM-DD 파싱, 실패 시 0 (now 미사용: 결정적)
package cluster

import "time"

// epochDay converts a front matter date scalar into days since the Unix
// epoch. It never consults the current time, so candidate ranking is
// deterministic. Unparseable or empty dates map to day 0.
func epochDay(date string) int64 {
	if t, err := time.Parse(time.RFC3339, date); err == nil {
		return t.Unix() / 86400
	}
	if t, err := time.Parse("2006-01-02", date); err == nil {
		return t.Unix() / 86400
	}
	return 0
}
