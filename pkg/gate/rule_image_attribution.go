//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [image-attribution] structure.order에 attribution이 선언된 경우 메인 이미지 직후 이탤릭 저작자 표기 검증
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleImageAttribution requires an italic attribution line (*Image: ...*)
// right after the main image. Without a main image it still fires, reporting
// the dependency (image-first names the root cause).
func ruleImageAttribution(t *Target) []blogyaml.Diagnostic {
	if !orderHas(t.Blog, "attribution") {
		return nil
	}
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if a.Doc.FirstIsImage && a.Doc.AttribLine >= 0 {
			continue
		}
		msg := "attribution line (*Image: ...*) must follow the main image, got none"
		if !a.Doc.FirstIsImage {
			msg = "no main image to attribute (see image-first)"
		}
		diags = append(diags, blogyaml.Diagnostic{
			File: a.Path, Line: bodyLine(a.Doc, a.Doc.FirstContentLine), Rule: "image-attribution",
			Message: msg,
		})
	}
	return diags
}
