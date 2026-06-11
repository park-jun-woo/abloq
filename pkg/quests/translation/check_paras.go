//ff:func feature=quest type=rule control=sequence topic=lossless
//ff:what 패리티 ② — fence-aware 블록 수의 원문↔번역 1:1 일치 검사 (문단 분할 보존)
package translation

import (
	"fmt"

	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// checkParas requires the translation to keep the origin's paragraph-block
// segmentation: same fence-aware block count.
func checkParas(where string, o, t *agate.Doc) []quest.Fact {
	oN, tN := fenceBlocks(o), fenceBlocks(t)
	if oN == tN {
		return nil
	}
	return []quest.Fact{{Where: where + "#paragraphs",
		Expected: fmt.Sprintf("%d paragraph block(s) as in the origin", oN),
		Actual:   fmt.Sprintf("%d block(s)", tN)}}
}
