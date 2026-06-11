//ff:func feature=quest type=rule control=iteration dimension=1
//ff:what abloq 게이트 카탈로그(gate.Rules)에서 룰ID의 설명 조회 — 어댑터 RuleMeta.Desc와 Fact.Expected의 단일 출처 (퀘스트 공용)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// RuleDesc looks one rule's description up in the abloq gate catalog so the
// adapter's rulebook entry and Fact.Expected never drift from the rule code.
func RuleDesc(id string) string {
	for _, r := range agate.Rules() {
		if r.ID == id {
			return r.Desc
		}
	}
	return ""
}
