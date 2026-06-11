//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 펜스 코드블록 전수 추출 — 여는 펜스 라인(언어 태그 포함)부터 내용까지를 블록 1개=문자열 1개로, 번역 금지 비교 대상 (패리티 ④)
package translation

import (
	"strings"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// codeBlocks extracts every fenced code block as one string: the opening
// fence line (its info/language tag included) plus the raw content lines,
// newline-joined, closing fence excluded. Code is translation-forbidden, so
// the parity rule compares the blocks as a bidirectional multiset.
func codeBlocks(d *agate.Doc) []string {
	var blocks []string
	var cur []string
	inFence := false
	for _, raw := range d.BodyLines {
		ln := strings.TrimSpace(raw)
		fence := strings.HasPrefix(ln, "```") || strings.HasPrefix(ln, "~~~")
		if fence && inFence {
			blocks = append(blocks, strings.Join(cur, "\n"))
			cur, inFence = nil, false
			continue
		}
		if !fence && !inFence {
			continue
		}
		cur = append(cur, raw)
		inFence = true
	}
	return blocks
}
