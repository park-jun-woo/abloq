//ff:func feature=visibility type=scorer control=sequence
//ff:what date 스칼라 → epoch 일수 점수 — RFC3339 또는 YYYY-MM-DD 파싱, 실패 시 0 (now 미사용: 결정적)
package priority

import "time"

// dateScore converts a front matter date scalar into days since the Unix
// epoch. It never consults the current time, so the score is deterministic.
// Unparseable or empty dates score 0 (lowest priority).
func dateScore(date string) int64 {
	if t, err := time.Parse(time.RFC3339, date); err == nil {
		return t.Unix() / 86400
	}
	if t, err := time.Parse("2006-01-02", date); err == nil {
		return t.Unix() / 86400
	}
	return 0
}
