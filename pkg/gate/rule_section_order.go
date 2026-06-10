//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [section-order] 인식된 섹션 헤딩이 structure.order의 상대 순서를 지키는지 글별 검증 — layout 특수 페이지는 스킵
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleSectionOrder requires recognized sections to appear in the canonical
// relative order declared by structure.order. Layout-owned special pages are
// exempt (see special).
func ruleSectionOrder(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if special(a) {
			continue
		}
		if d := sectionOrderDiag(t.heads.rank, a); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}
