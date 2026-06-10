//ff:func feature=gate type=rule control=iteration dimension=1 topic=lossless
//ff:what [body-lossless] 원본(git HEAD) 본문 라인 multiset이 현재본에 모두 보존되는지 검증 — 본문 삭제·변조 치즈 방어
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleBodyLossless requires every baseline body line (excluding structural
// lines) to survive in the current document, as a multiset subset.
func ruleBodyLossless(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if a.Base == nil || a.Base == a.Doc {
			continue
		}
		miss, ok := MultisetSubset(ContentLines(a.Base), ContentLines(a.Doc))
		if ok {
			continue
		}
		diags = append(diags, blogyaml.Diagnostic{
			File: a.Path, Line: 1, Rule: "body-lossless",
			Message: "baseline body line deleted or altered: " + trunc(miss),
		})
	}
	return diags
}
