//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [image-first] structure.order에 image가 선언된 경우 첫 비공백 본문 라인이 메인 이미지인지 검증 — layout 특수 페이지는 스킵
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleImageFirst requires the first content line to be the main image.
// Layout-owned special pages are exempt (see special).
func ruleImageFirst(t *Target) []blogyaml.Diagnostic {
	if !orderHas(t.Blog, "image") {
		return nil
	}
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if a.Doc.FirstIsImage || special(a) {
			continue
		}
		actual := "<no content>"
		if a.Doc.FirstContentLine >= 0 {
			actual = trunc(a.Doc.BodyLines[a.Doc.FirstContentLine])
		}
		diags = append(diags, blogyaml.Diagnostic{
			File: a.Path, Line: bodyLine(a.Doc, a.Doc.FirstContentLine), Rule: "image-first",
			Message: "first content line must be the main image ![...](...), got " + actual,
		})
	}
	return diags
}
