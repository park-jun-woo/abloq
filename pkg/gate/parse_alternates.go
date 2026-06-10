//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what HTML head의 rel=alternate 링크에서 hreflang→href 맵 추출
package gate

import "regexp"

var reAlternate = regexp.MustCompile(`<link[^>]*\brel="alternate"[^>]*\bhreflang="([^"]+)"[^>]*\bhref="([^"]*)"`)

// parseAlternates extracts the hreflang alternate links of a built page.
func parseAlternates(html string) map[string]string {
	alts := map[string]string{}
	for _, m := range reAlternate.FindAllStringSubmatch(html, -1) {
		alts[m[1]] = m[2]
	}
	return alts
}
