//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what [numeric-claim-sourced] 글마다 무출처 수치 주장을 진단 — 수치+단위+단정 검출, 같은 문단 출처 링크 요구
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleNumericClaimSourced fails every numeric claim (number+unit+assertion)
// whose paragraph carries no source link. Code and quote blocks are excluded
// by the detector; claims_ignore (with reasons) exempts an article.
func ruleNumericClaimSourced(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		diags = append(diags, numClaimDiags(a)...)
	}
	return diags
}
