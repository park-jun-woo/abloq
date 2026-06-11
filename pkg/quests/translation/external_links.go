//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 링크 목적지에서 외부 URL(http/https)만 선별 — 번역 불변 비교 대상 (패리티 ⑤)
package translation

import (
	"strings"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// externalLinks returns the http(s) link destinations in the body prose.
// External URLs are translation-invariant; the parity rule compares them as a
// bidirectional multiset (injection is caught, not just loss).
func externalLinks(d *agate.Doc) []string {
	var urls []string
	for _, dest := range linkDests(d) {
		if !strings.HasPrefix(dest, "http://") && !strings.HasPrefix(dest, "https://") {
			continue
		}
		urls = append(urls, dest)
	}
	return urls
}
