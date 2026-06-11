//ff:func feature=insight type=rule control=sequence
//ff:what 검증 룰 4종(claims-min/claim-id-unique/claim-kind/claim-anchors-empty)을 순서대로 실행 — (에러, 경고) 분리 반환
package insight

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Validate runs all insight.yaml validation rules. Anchors-empty findings are
// warnings (the claim cannot be screened by match) and never block.
func Validate(filename string, ins *Insight, lines []int) (errs, warns []blogyaml.Diagnostic) {
	errs = append(errs, ruleClaimsMin(filename, ins)...)
	errs = append(errs, ruleClaimIDUnique(filename, ins, lines)...)
	errs = append(errs, ruleClaimKind(filename, ins, lines)...)
	warns = append(warns, ruleAnchorsEmpty(filename, ins, lines)...)
	return errs, warns
}
