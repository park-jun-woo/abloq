//ff:func feature=quest type=rule control=sequence
//ff:what abloq 게이트 카탈로그에서 룰ID의 설명 조회 — 공용 추출본(quests/common.RuleDesc) 위임
//ff:why Phase017에서 번역 퀘스트와 공유하도록 구현을 pkg/quests/common으로 추출 — 복제 금지, writing은 추출본을 쓴다
package writing

import "github.com/park-jun-woo/abloq/pkg/quests/common"

// ruleDesc looks one rule's description up in the abloq gate catalog. The
// implementation lives in pkg/quests/common (Phase017 extraction).
func ruleDesc(id string) string {
	return common.RuleDesc(id)
}
