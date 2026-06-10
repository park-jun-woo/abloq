//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what 언어 1개에 대해 blog.yaml 선언 섹션을 순서대로 순회해 대상 글 수집
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// discoverLang collects articles for one language across declared sections.
func discoverLang(dir string, b *blogyaml.Blog, hi headingIndex, lang string) []*Article {
	var arts []*Article
	for _, section := range b.Sections {
		arts = append(arts, discoverSection(dir, hi, lang, section)...)
	}
	return arts
}
