//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [slug-consistency] 같은 글(섹션/파일 어간)의 전 언어 slug 일치와 누락 언어를 그룹 단위로 검출
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ruleSlugConsistency groups articles by section/file-stem and checks that
// every declared language has the article and that effective slugs agree.
func ruleSlugConsistency(t *Target) []blogyaml.Diagnostic {
	groups := groupArticles(t.Articles)
	var diags []blogyaml.Diagnostic
	for _, key := range sortedGroupKeys(groups) {
		diags = append(diags, missingLangDiags(t.Blog.Languages, key, groups[key])...)
		if d := slugMismatchDiag(groups[key]); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}
