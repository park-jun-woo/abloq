//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [section-order] 인식된 섹션 헤딩이 structure.order의 상대 순서를 지키는지 글별 검증
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleSectionOrder requires recognized sections to appear in the canonical
// relative order declared by structure.order.
func ruleSectionOrder(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if d := sectionOrderDiag(t.heads.rank, a); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}
