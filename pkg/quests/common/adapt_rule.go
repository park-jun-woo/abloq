//ff:func feature=quest type=rule control=sequence
//ff:what abloq 게이트 룰 1개를 reins 룰로 감싸는 공용 어댑터 — gate.Run(단일 룰)을 호출해 발동 진단을 Fact 1건으로 매핑, LevelFail 고정
//ff:why Phase016 writing 프라이빗 구현을 Phase017에서 공용 추출 — 퀘스트마다 복제하지 않고 TargetCarrier 계약으로 제출물 타입에 비의존
package common

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// AdaptRule wraps one abloq gate rule (by catalog ID) as a reins LevelFail
// rule: Check runs exactly that rule over the submission's target and maps
// its diagnostics to a single located Fact (existing rule code unchanged).
func AdaptRule(id string) rgate.Rule {
	desc := RuleDesc(id)
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: id, Level: rgate.LevelFail, Desc: desc},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(TargetCarrier)
			diags := agate.Run(sub.GateTarget(), id)
			if len(diags) == 0 {
				return false, quest.Fact{}
			}
			return true, DiagsFact(desc, diags)
		},
	}
}
