//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what insight-claim-id-unique — claims[].id 중복(빈 id 포함)을 에러로 진단
package insight

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// ruleClaimIDUnique flags duplicate (or duplicated-empty) claim ids.
func ruleClaimIDUnique(filename string, ins *Insight, lines []int) []blogyaml.Diagnostic {
	seen := make(map[string]bool, len(ins.Claims))
	var diags []blogyaml.Diagnostic
	for i, c := range ins.Claims {
		if seen[c.ID] {
			diags = append(diags, blogyaml.Diagnostic{
				File: filename, Line: claimLine(lines, i), Rule: "insight-claim-id-unique",
				Message: fmt.Sprintf("duplicate claim id %q", c.ID),
			})
			continue
		}
		seen[c.ID] = true
	}
	return diags
}
