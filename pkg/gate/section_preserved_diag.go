//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what 글 1편에서 원본 대비 사라진 첫 인식 섹션 키를 찾아 진단으로 반환
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// sectionPreservedDiag reports the first baseline section key missing from the
// current document.
func sectionPreservedDiag(a *Article) *blogyaml.Diagnostic {
	have := presentKeys(a.Doc.Sections)
	for _, s := range a.Base.Sections {
		if have[s.Key] {
			continue
		}
		return &blogyaml.Diagnostic{
			File: a.Path, Line: 1, Rule: "section-preserved",
			Message: "section " + s.Key + " (" + trunc(s.Text) + ") present in baseline was removed",
		}
	}
	return nil
}
