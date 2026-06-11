//ff:func feature=quest type=rule control=sequence
//ff:what abloq 진단 목록 → reins Fact 1건 — 첫 진단의 파일:라인/메시지, 다중 진단은 "(외 N건)" 병기 (룰당 Fact 1건 규약)
package writing

import (
	"fmt"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// diagsFact maps one fired rule's diagnostics to the single Fact the adapter
// emits: the first diagnostic located and quantified, with the remaining
// count noted. The Rule field is stamped by reins gate.Evaluate.
func diagsFact(expected string, diags []blogyaml.Diagnostic) quest.Fact {
	d := diags[0]
	actual := d.Message
	if len(diags) > 1 {
		actual += fmt.Sprintf(" (외 %d건)", len(diags)-1)
	}
	return quest.Fact{
		Where:    fmt.Sprintf("%s:%d", d.File, d.Line),
		Expected: expected,
		Actual:   actual,
	}
}
