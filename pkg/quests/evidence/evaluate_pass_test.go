//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what 게이트 전 룰 통과 검증 — 큐 주장 출처 추가+rot 교체(큐 밖 주장 불변·lastmod 불변경)의 정상 보강을 reins 레벨집계로 평가해 PASS
package evidence

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluatePass(t *testing.T) {
	root := writeInstance(t)
	sourced := strings.Replace(baseArticleMD,
		unsourcedClaim,
		unsourcedClaim+" [Migration report](https://example.org/spec)", 1)
	sourced = strings.Replace(sourced, rotURL, "https://example.org/live-study", 1)
	writeFile(t, root, "content/en/posts/a.md", sourced)
	v := rgate.Evaluate(Definition{}.Rules(), subWith(t, root))
	if v.Outcome != quest.OutPass {
		t.Fatalf("Outcome = %s, facts = %+v — want PASS", v.Outcome, v.Facts)
	}
}
