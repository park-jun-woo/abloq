//ff:func feature=content type=parser control=iteration dimension=1
//ff:what sources 헤딩 아래 리스트 항목(-, *, n.) 수 집계 — 헤딩이 없으면 0 (gate min-sources와 같은 집계 규칙)
package content

import (
	"regexp"
	"strings"
)

var reOrderedItem = regexp.MustCompile(`^\d+\. `)

// sourceCount counts markdown list items between the localized sources
// heading and the next heading. heading is the blog.yaml
// structure.headings.sources text for the article's language; empty or absent
// headings yield 0.
func sourceCount(body, heading string) int64 {
	if heading == "" {
		return 0
	}
	var n int64
	in := false
	for _, raw := range strings.Split(body, "\n") {
		ln := strings.TrimSpace(raw)
		if strings.HasPrefix(ln, "#") {
			in = strings.TrimSpace(strings.TrimLeft(ln, "#")) == heading
			continue
		}
		if in && (strings.HasPrefix(ln, "- ") || strings.HasPrefix(ln, "* ") || reOrderedItem.MatchString(ln)) {
			n++
		}
	}
	return n
}
