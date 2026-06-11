//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what 본문 전체 헤딩 레벨 시퀀스 추출(H1~H6, 코드 펜스 밖) — 인식 섹션만 보는 headingIndex와 달리 자유 헤딩 포함 (패리티 ①)
package translation

import (
	"regexp"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

var reAnyHeading = regexp.MustCompile(`^(#{1,6})\s+\S`)

// headingLevels returns the document-order sequence of every heading level in
// the body prose — free-form H2/H3 included, unlike the gate's headingIndex
// which only recognizes blog.yaml-declared section headings.
func headingLevels(d *agate.Doc) []int {
	var levels []int
	for _, raw := range proseLines(d) {
		m := reAnyHeading.FindStringSubmatch(raw)
		if m == nil {
			continue
		}
		levels = append(levels, len(m[1]))
	}
	return levels
}
