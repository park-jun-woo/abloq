//ff:func feature=gen type=generator control=selection
//ff:what llms.txt 언어 스코프 결정 — base(기본 언어 1개, 기본값)/all(전 언어)/명시 목록을 실제 언어 리스트로 환원
package llms

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// scopeLangs resolves geo.llms_txt.languages into the concrete language list:
// "all" -> every declared language, "base" (or unset) -> the first declared
// language only, otherwise the explicit list as declared.
func scopeLangs(b *blogyaml.Blog) []string {
	spec := b.Geo.LlmsTxt.Languages
	switch {
	case len(spec) == 1 && spec[0] == "all":
		return b.Languages
	case len(spec) == 0 || (len(spec) == 1 && spec[0] == "base"):
		return b.Languages[:min(1, len(b.Languages))]
	default:
		return spec
	}
}
