//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what [claim-scope] 기준선의 수치 주장 중 큐 payload 해시 밖의 것이 Doc에 그대로(텍스트 multiset) 보존됐는지 검증 — 큐에 없는 주장 변경은 FAIL
//ff:why "큐에 없는 주장 변경 금지"의 결정적 구현 — 근거 보강을 빙자해 무관한 주장을 리워딩하는 치즈를 막는다. 해시는 Phase010 hashText(gate.HashText) 그대로라 payload claims와 1:1 대조된다 (Phase018 계획)
package common

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// RuleClaimScope builds the shared claim-scope rule: every baseline numeric
// claim whose hash is NOT in the queue payload must survive verbatim in the
// submitted article (text multiset — line-level, so claim lines must not be
// rewrapped). Claims the queue authorizes (QueuedClaims) may change freely.
func RuleClaimScope() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "claim-scope", Level: rgate.LevelFail,
			Desc: "numeric claims outside the queued claim hashes are preserved verbatim"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			c := ctx.Submission.(ConsCarrier).Cons()
			a := c.Target.Articles[0]
			if a.Base == nil {
				return false, quest.Fact{}
			}
			want := claimTexts(a.Base, c.QueuedClaims)
			have := claimTexts(a.Doc, nil)
			missing, ok := agate.MultisetSubset(want, have)
			if ok {
				return false, quest.Fact{}
			}
			return true, quest.Fact{Where: a.Path,
				Expected: "claim line preserved verbatim (only claims listed in the queue payload may change)",
				Actual:   "baseline claim missing or altered: " + missing}
		},
	}
}
