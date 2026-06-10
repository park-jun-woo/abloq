//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [front-matter-schema] 글별 front matter 필수 필드(title/date/lastmod/tags)와 타입 검증 — layout 특수 페이지는 스킵
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleFrontMatterSchema validates each article's front matter fields.
// Layout-owned special pages declare their own front matter contract with
// their layout, so the article schema does not apply (see special).
func ruleFrontMatterSchema(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if special(a) {
			continue
		}
		diags = append(diags, fmSchemaDiags(a)...)
	}
	return diags
}
