//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what sources 섹션의 출처 항목 수 집계 — 리스트 항목(-, *, n.)만 출처로 인정, 섹션 없으면 false
package gate

import (
	"regexp"
	"strings"
)

var reOrderedItem = regexp.MustCompile(`^\d+\. `)

// sourceCount counts the source entries (markdown list items) inside the
// sources section. headLine is the heading's BodyLines index; found is false
// when the article has no recognized sources section.
func sourceCount(d *Doc) (n, headLine int, found bool) {
	start, end, ok := sectionSpan(d, "sources")
	if !ok {
		return 0, 0, false
	}
	for _, raw := range d.BodyLines[start+1 : end] {
		ln := strings.TrimSpace(raw)
		if strings.HasPrefix(ln, "- ") || strings.HasPrefix(ln, "* ") || reOrderedItem.MatchString(ln) {
			n++
		}
	}
	return n, start, true
}
