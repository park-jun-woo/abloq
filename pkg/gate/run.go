//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what 게이트 실행기 — 등록된 룰 전부(또는 지정 룰ID만)를 순서대로 실행해 진단을 모음
package gate

import (
	"slices"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Run executes the registered rules against the target. With ruleIDs given,
// only those rules run; with none, all rules run.
func Run(t *Target, ruleIDs ...string) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, r := range Rules() {
		if len(ruleIDs) > 0 && !slices.Contains(ruleIDs, r.ID) {
			continue
		}
		diags = append(diags, r.Check(t)...)
	}
	return diags
}
