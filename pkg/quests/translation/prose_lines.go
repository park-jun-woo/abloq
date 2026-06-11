//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 코드 펜스 밖 본문 라인만 추출 — 이미지/링크/헤딩 추출기가 펜스 안 텍스트를 오인하지 않게 하는 공통 전처리
package translation

import (
	"strings"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// proseLines returns the body lines outside fenced code blocks (fence
// delimiter lines excluded). Every prose-level extractor (headings, images,
// links) runs over these so fence contents never masquerade as markup.
func proseLines(d *agate.Doc) []string {
	var out []string
	inFence := false
	for _, raw := range d.BodyLines {
		ln := strings.TrimSpace(raw)
		if strings.HasPrefix(ln, "```") || strings.HasPrefix(ln, "~~~") {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		out = append(out, raw)
	}
	return out
}
