//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what [honest-lastmod] lastmod 갱신 글의 본문 실변경(정규화 토큰 diff ≥ 임계)과 신선도 큐 등재를 검증
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleHonestLastmod blocks lastmod forgery: an updated lastmod requires a
// meaningful body change and, when a freshness queue exists, queue membership.
func ruleHonestLastmod(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		diags = append(diags, honestLastmodDiags(t, a)...)
	}
	return diags
}
