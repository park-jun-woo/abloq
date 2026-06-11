//ff:func feature=quest type=rule control=iteration dimension=1 topic=queue
//ff:what [queue-scope] 작업트리 변경 파일 집합 ⊆ 허용 집합(대상 글+전 언어 번역본+insight 사이드카+kind별 추가) 검증 — 범위 밖 변경은 FAIL
//ff:why 산문 노동을 큐 payload 범위 안에 가둔다 — 다른 글·blog.yaml·레이아웃을 고쳐 게이트를 우회하거나 큐 파일을 미리 지우는 치즈를 변경 집합 포함 관계 하나로 차단한다 (Phase018 계획)
package common

import (
	"fmt"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// RuleQueueScope builds the shared queue-scope rule: every working-tree
// change captured at Prepare time must fall inside the item's allowed file
// set. The queue file itself is never allowed — its deletion is the
// post-gate consumption signal, not part of the submission.
func RuleQueueScope() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "queue-scope", Level: rgate.LevelFail,
			Desc: "every working-tree change stays inside the queue item's allowed file set"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			c := ctx.Submission.(ConsCarrier).Cons()
			var off []string
			for _, p := range c.Changed {
				if !c.Allowed[p] {
					off = append(off, p)
				}
			}
			if len(off) == 0 {
				return false, quest.Fact{}
			}
			actual := "out-of-scope change: " + off[0]
			if len(off) > 1 {
				actual += fmt.Sprintf(" (외 %d건)", len(off)-1)
			}
			return true, quest.Fact{Where: off[0],
				Expected: "changes limited to the target article, its language companions and insight sidecars",
				Actual:   actual}
		},
	}
}
