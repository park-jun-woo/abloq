//ff:func feature=quest type=rule control=sequence
//ff:what abloq 진단 목록 → reins Fact 1건 — 공용 추출본(quests/common.DiagsFact) 위임
//ff:why Phase017에서 번역 퀘스트와 공유하도록 구현을 pkg/quests/common으로 추출 — 복제 금지, writing은 추출본을 쓴다
package writing

import (
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// diagsFact maps one fired rule's diagnostics to the single Fact the adapter
// emits. The implementation lives in pkg/quests/common (Phase017 extraction).
func diagsFact(expected string, diags []blogyaml.Diagnostic) quest.Fact {
	return common.DiagsFact(expected, diags)
}
