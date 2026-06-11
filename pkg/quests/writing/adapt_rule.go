//ff:func feature=quest type=rule control=sequence
//ff:what abloq 게이트 룰 1개를 reins 룰로 감싸는 어댑터 — gate.Run(단일 룰)을 호출해 발동 진단을 Fact 1건으로 매핑, LevelFail 고정
package writing

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// adaptRule wraps one abloq gate rule (by catalog ID) as a reins LevelFail
// rule: Check runs exactly that rule over the submission's target and maps
// its diagnostics to a single located Fact (existing rule code unchanged).
func adaptRule(id string) rgate.Rule {
	desc := ruleDesc(id)
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: id, Level: rgate.LevelFail, Desc: desc},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			diags := agate.Run(sub.Target, id)
			if len(diags) == 0 {
				return false, quest.Fact{}
			}
			return true, diagsFact(desc, diags)
		},
	}
}
