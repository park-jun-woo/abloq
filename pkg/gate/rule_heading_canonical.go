//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [heading-canonical] 이름으로 인식된 섹션 헤딩이 ## 레벨이 아닌 글을 검출 — 모호 케이스는 사람 판단 대상
package gate

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// ruleHeadingCanonical flags section names found at a non-## heading level.
func ruleHeadingCanonical(t *Target) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		if len(a.Doc.BadLevel) == 0 {
			continue
		}
		h := a.Doc.BadLevel[0]
		diags = append(diags, blogyaml.Diagnostic{
			File: a.Path, Line: bodyLine(a.Doc, h.Line), Rule: "heading-canonical",
			Message: fmt.Sprintf("section %s heading must be ## level, got %d-level %q", h.Key, h.Level, trunc(h.Text)),
		})
	}
	return diags
}
