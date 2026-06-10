//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what [section-preserved] 원본(git HEAD)에 있던 인식 섹션이 현재본에서도 유지되는지 검증 — 섹션 삭제 치즈 방어
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleSectionPreserved blocks the "delete a section to fix the order" cheese.
// Articles without a distinct baseline are skipped.
func ruleSectionPreserved(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if a.Base == nil || a.Base == a.Doc {
			continue
		}
		if d := sectionPreservedDiag(a); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}
