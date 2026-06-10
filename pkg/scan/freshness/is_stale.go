//ff:func feature=scan type=rule control=sequence
//ff:what 신선도 판정 — now - lastmod > freshness_days, lastmod 미파싱이면 판정 불가로 false
package freshness

import "time"

// isStale reports whether the article's lastmod exceeded the freshness
// window. An unparseable lastmod cannot be judged and never queues.
func isStale(lastmod string, days int, now time.Time) bool {
	t, ok := parseDate(lastmod)
	if !ok {
		return false
	}
	return now.Sub(t) > time.Duration(days)*24*time.Hour
}
