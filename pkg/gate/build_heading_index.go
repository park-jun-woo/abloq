//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what blog.yaml structure에서 헤딩 인덱스 구축 — order 순위 + 언어별 정규화 헤딩 역색인
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// buildHeadingIndex derives the heading lookup from blog.yaml structure.
// Matching is by normalized text (normText), never by stem.
func buildHeadingIndex(b *blogyaml.Blog) headingIndex {
	hi := headingIndex{byLang: map[string]map[string]string{}, rank: map[string]int{}}
	for i, key := range b.Structure.Order {
		hi.rank[key] = i
	}
	for key, langs := range b.Structure.Headings {
		indexHeadingKey(hi.byLang, key, langs)
	}
	return hi
}
