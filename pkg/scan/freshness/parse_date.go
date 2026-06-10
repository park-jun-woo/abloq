//ff:func feature=scan type=parser control=sequence
//ff:what front matter 날짜 스칼라 파싱 — RFC3339 또는 YYYY-MM-DD, 실패 시 ok=false
package freshness

import "time"

// parseDate parses a front matter date/lastmod scalar. The two accepted
// layouts mirror what the indexer passes through verbatim.
func parseDate(s string) (time.Time, bool) {
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, true
	}
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t, true
	}
	return time.Time{}, false
}
