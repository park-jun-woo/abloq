//ff:func feature=content type=parser control=iteration dimension=1
//ff:what 언어 1개에 대해 blog.yaml 선언 섹션을 순서대로 순회해 인덱스 항목 수집
package content

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// indexLang gathers index entries for one language across declared sections.
func indexLang(root string, b *blogyaml.Blog, lang string) []Entry {
	var entries []Entry
	for _, section := range b.Sections {
		entries = append(entries, indexSection(root, b, lang, section)...)
	}
	return entries
}
