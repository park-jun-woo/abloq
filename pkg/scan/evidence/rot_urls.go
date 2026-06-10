//ff:func feature=scan type=rule control=iteration dimension=1 topic=evidence
//ff:what 글 1편의 확정 rot URL 목록 — 연속 실패 rotThreshold(3)회 이상만, 본문 등장 순서 유지
package evidence

// rotThreshold is how many consecutive failed scans confirm link rot —
// transient 404s and CDN hiccups never persist this long at scan cadence.
const rotThreshold = 3

// rotURLs filters one article's confirmed-rot citations from the updated
// check state.
func rotURLs(checks []Check, lang, section, slug string) []string {
	var urls []string
	for _, c := range checks {
		if c.Lang != lang || c.Section != section || c.Slug != slug || c.ConsecutiveFailures < rotThreshold {
			continue
		}
		urls = append(urls, c.URL)
	}
	return urls
}
