//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what insight-claim-kind — claims[].kind가 claim|rebuttal|prediction|definition 밖이면 에러
package insight

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// claimKinds is the closed kind vocabulary (주장|반론 대응|예측|정의).
var claimKinds = map[string]bool{"claim": true, "rebuttal": true, "prediction": true, "definition": true}

// ruleClaimKind flags claims whose kind is outside the closed vocabulary.
func ruleClaimKind(filename string, ins *Insight, lines []int) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for i, c := range ins.Claims {
		if claimKinds[c.Kind] {
			continue
		}
		diags = append(diags, blogyaml.Diagnostic{
			File: filename, Line: claimLine(lines, i), Rule: "insight-claim-kind",
			Message: fmt.Sprintf("claim %q: kind must be claim|rebuttal|prediction|definition, got %q", c.ID, c.Kind),
		})
	}
	return diags
}
