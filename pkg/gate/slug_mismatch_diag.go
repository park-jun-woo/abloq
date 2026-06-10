//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 글 그룹 1개의 유효 slug(front matter slug 또는 파일 어간)가 전 언어에서 일치하는지 검사
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// slugMismatchDiag reports the first language version whose effective slug
// differs from the group's first one.
func slugMismatchDiag(group []*Article) *blogyaml.Diagnostic {
	want := effSlug(group[0])
	for _, a := range group[1:] {
		got := effSlug(a)
		if got == want {
			continue
		}
		return &blogyaml.Diagnostic{
			File: a.Path, Line: fmKeyLine(a.Doc.FrontMatter, "slug"), Rule: "slug-consistency",
			Message: "slug " + got + " differs from " + want + " (" + group[0].Lang + " version)",
		}
	}
	return nil
}
