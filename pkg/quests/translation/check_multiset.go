//ff:func feature=quest type=rule control=sequence topic=lossless
//ff:what 양방향 multiset 일치 검사 — 원문→번역 누락과 번역→원문 주입을 각각 Fact로 (패리티 ③④⑤⑥(b) 공통)
//ff:why MultisetSubset은 단방향 포함 비교라 한 번이면 주입(번역에만 있는 원소)이 통과한다 — 양방향 적용으로 누락과 주입을 모두 차단 (Phase017 계획)
package translation

import (
	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// checkMultiset compares one translation-invariant element multiset in both
// directions: every origin element must survive into the translation, and the
// translation must not inject elements absent from the origin.
func checkMultiset(where, what string, origin, trans []string) []quest.Fact {
	var facts []quest.Fact
	if miss, ok := agate.MultisetSubset(origin, trans); !ok {
		facts = append(facts, quest.Fact{Where: where,
			Expected: what + " preserved from the origin",
			Actual:   "missing in translation: " + miss})
	}
	if extra, ok := agate.MultisetSubset(trans, origin); !ok {
		facts = append(facts, quest.Fact{Where: where,
			Expected: "no " + what + " beyond the origin",
			Actual:   "injected in translation: " + extra})
	}
	return facts
}
