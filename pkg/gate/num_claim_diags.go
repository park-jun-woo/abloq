//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what 글 1편의 신규 무출처 수치 주장 진단 — claims_ignore(사유 필수) 예외 처리 후 git HEAD 대비 신규(Sourced=false) 주장마다 1건
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// numClaimDiags judges one article: a valid claims_ignore (reasons stated)
// exempts the whole article; otherwise every unsourced claim added since the
// git HEAD baseline is a violation (pre-existing claims are the scanner's job).
func numClaimDiags(a *Article) []blogyaml.Diagnostic {
	exempt, bad := claimsIgnore(a)
	if bad {
		return []blogyaml.Diagnostic{{File: a.Path, Line: fmKeyLine(a.Doc.FrontMatter, "claims_ignore"),
			Rule:    "numeric-claim-sourced",
			Message: "claims_ignore must be a non-empty list of reason strings (a reason is required for the exemption)"}}
	}
	if exempt {
		return nil
	}
	var diags []blogyaml.Diagnostic
	for _, c := range newClaims(a) {
		if c.Sourced {
			continue
		}
		diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: c.Line, Rule: "numeric-claim-sourced",
			Message: "numeric claim has no source link in its paragraph: " + trunc(c.Text)})
	}
	return diags
}
