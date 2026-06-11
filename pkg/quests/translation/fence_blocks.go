//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what fence-aware 블록 splitter — 코드 펜스 밖 빈 줄만 블록 경계로 세는 문단 카운터 (패리티 ②)
//ff:why 신설 근거: 기존 claimParas는 인용·코드·헤딩·이미지·sources 섹션을 제외하는 "주장 적격" 필터라(목록은 제외하지 않음) 레이아웃 문단 수 비교에 부적합 — 패리티는 모든 블록(인용·헤딩·이미지·목록 포함)을 세되 코드블록 내 빈 줄에 오염되지 않아야 한다 (Phase017 계획 ②)
package translation

import (
	"strings"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// fenceBlocks counts the body's paragraph-level blocks: maximal runs of
// non-blank lines, where blank lines inside fenced code never split a block.
// Paragraph segmentation is a layout property the translator controls, so the
// count is enforceable across CJK and RTL languages alike.
func fenceBlocks(d *agate.Doc) int {
	n, inBlock, inFence := 0, false, false
	for _, raw := range d.BodyLines {
		ln := strings.TrimSpace(raw)
		if strings.HasPrefix(ln, "```") || strings.HasPrefix(ln, "~~~") {
			inFence = !inFence
		}
		if ln == "" && !inFence {
			inBlock = false
			continue
		}
		if !inBlock {
			n++
			inBlock = true
		}
	}
	return n
}
