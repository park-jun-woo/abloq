//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what 인용 표기와 페이지 title/og:title의 토큰 겹침이 임계(0.3) 이상인지 판정 — 표기에 토큰이 없으면 통과
package gate

// citationMetaThreshold is the minimum token-overlap ratio (in either
// direction) between the citation label and the page title.
const citationMetaThreshold = 0.3

// metaMatch reports whether the cited page's title/og:title plausibly matches
// the citation label. A label with no tokens (e.g. a bare URL) claims nothing
// about the page, so it passes.
func metaMatch(html, label string) bool {
	want := Tokens(label)
	if len(want) == 0 {
		return true
	}
	for _, title := range htmlTitles(html) {
		got := Tokens(title)
		if tokenOverlap(want, got) >= citationMetaThreshold || tokenOverlap(got, want) >= citationMetaThreshold {
			return true
		}
	}
	return false
}
