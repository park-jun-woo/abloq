//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what insight-claim-anchors-empty — anchors가 빈 claim은 (requires_source 무관) match가 스크리닝할 수 없다는 경고
package insight

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// ruleAnchorsEmpty warns on claims that have no anchors: match cannot screen
// them, regardless of requires_source.
func ruleAnchorsEmpty(filename string, ins *Insight, lines []int) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for i, c := range ins.Claims {
		if len(c.Anchors) > 0 {
			continue
		}
		diags = append(diags, blogyaml.Diagnostic{
			File: filename, Line: claimLine(lines, i), Rule: "insight-claim-anchors-empty",
			Message: fmt.Sprintf("claim %q has no anchors — match cannot screen it (warning)", c.ID),
		})
	}
	return diags
}
