//ff:func feature=insight type=rule control=sequence
//ff:what insight-claims-min — claims가 1개 미만이면 에러 (대조할 주장이 없으면 게이트가 설 수 없다)
package insight

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleClaimsMin requires at least one claim.
func ruleClaimsMin(filename string, ins *Insight) []blogyaml.Diagnostic {
	if len(ins.Claims) > 0 {
		return nil
	}
	return []blogyaml.Diagnostic{{
		File: filename, Line: 1, Rule: "insight-claims-min",
		Message: "claims must contain at least 1 claim",
	}}
}
