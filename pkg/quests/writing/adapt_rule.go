//ff:func feature=quest type=rule control=sequence
//ff:what abloq 게이트 룰 1개를 reins 룰로 감싸는 어댑터 — 공용 추출본(quests/common.AdaptRule) 위임
//ff:why Phase017에서 번역 퀘스트와 공유하도록 구현을 pkg/quests/common으로 추출 — 복제 금지, writing은 추출본을 쓴다
package writing

import (
	rgate "github.com/park-jun-woo/reins/pkg/gate"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// adaptRule wraps one abloq gate rule (by catalog ID) as a reins LevelFail
// rule. The implementation lives in pkg/quests/common (Phase017 extraction);
// Submission satisfies common.TargetCarrier.
func adaptRule(id string) rgate.Rule {
	return common.AdaptRule(id)
}
