//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 글 그룹 1개에서 선언 언어 중 해당 글이 없는 언어를 찾아 언어별 진단 생성
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// missingLangDiags reports each declared language missing from one article
// group.
func missingLangDiags(langs []string, key string, group []*Article) []blogyaml.Diagnostic {
	have := map[string]bool{}
	for _, a := range group {
		have[a.Lang] = true
	}
	var diags []blogyaml.Diagnostic
	for _, lang := range langs {
		if have[lang] {
			continue
		}
		diags = append(diags, blogyaml.Diagnostic{
			File: group[0].Path, Line: 1, Rule: "slug-consistency",
			Message: "article " + key + " has no " + lang + " version",
		})
	}
	return diags
}
