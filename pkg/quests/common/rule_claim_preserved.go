//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what [claim-preserved] Doc의 수치 주장 건수 ≥ 기준선 건수(글 단위) 검증 — 감소는 FAIL
//ff:why 갱신은 낡은 수치의 교체가 본질이라 텍스트 보존(claim-scope)은 모순 — "기존 주장 삭제 금지"의 결정적 프록시가 건수 하한이다. 의미 수준 보존(주장 대체의 타당성)은 비결정이라 게이트 룰이 아니고 REVIEW 안내로 남긴다 (Phase018 계획)
package common

import (
	"fmt"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// RuleClaimPreserved builds the shared claim-preserved rule: the submitted
// article must carry at least as many numeric claims as its baseline.
// Refreshing replaces stale figures, so verbatim preservation cannot apply —
// the count floor is the deterministic proxy for "never delete claims".
func RuleClaimPreserved() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "claim-preserved", Level: rgate.LevelFail,
			Desc: "the article keeps at least the baseline's numeric claim count"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			c := ctx.Submission.(ConsCarrier).Cons()
			a := c.Target.Articles[0]
			if a.Base == nil {
				return false, quest.Fact{}
			}
			base := len(agate.DetectClaims(a.Base))
			doc := len(agate.DetectClaims(a.Doc))
			if doc >= base {
				return false, quest.Fact{}
			}
			return true, quest.Fact{Where: a.Path,
				Expected: fmt.Sprintf(">= %d numeric claim(s) (baseline count)", base),
				Actual:   fmt.Sprintf("%d claim(s) — refreshing must replace stale figures, not delete them", doc)}
		},
	}
}
