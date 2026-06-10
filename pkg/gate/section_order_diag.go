//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 글 1편의 인식 섹션열에서 첫 순서 역전(앞 섹션의 order 순위 > 뒤 섹션)을 찾아 진단으로 반환
package gate

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// sectionOrderDiag returns the first out-of-order section pair of one article.
func sectionOrderDiag(rank map[string]int, a *Article) *blogyaml.Diagnostic {
	secs := a.Doc.Sections
	for i := 1; i < len(secs); i++ {
		if rank[secs[i].Key] >= rank[secs[i-1].Key] {
			continue
		}
		return &blogyaml.Diagnostic{
			File: a.Path, Line: bodyLine(a.Doc, secs[i].Line), Rule: "section-order",
			Message: fmt.Sprintf("section %q(%s) must come before %q(%s)",
				secs[i].Text, secs[i].Key, secs[i-1].Text, secs[i-1].Key),
		}
	}
	return nil
}
