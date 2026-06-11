//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what 게이트 전 룰 통과 검증 — 대상 글 out 링크+후보 글 in 앵커(lastmod 불변경)의 정상 큐레이션을 reins 레벨집계로 평가해 PASS
package cluster

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluatePass(t *testing.T) {
	root := writeInstance(t)
	writeFile(t, root, "content/en/posts/thin.md",
		thinArticleMD+"\nSee the [hub](/posts/hub/) overview.\n")
	writeFile(t, root, "content/en/posts/hub.md",
		hubArticleMD+"\nThe [thin](/posts/thin/) article covers the edge case.\n")
	v := rgate.Evaluate(Definition{}.Rules(), subWith(t, root))
	if v.Outcome != quest.OutPass {
		t.Fatalf("Outcome = %s, facts = %+v — want PASS", v.Outcome, v.Facts)
	}
}
