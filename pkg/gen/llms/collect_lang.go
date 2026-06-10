//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what 언어 1개에 대해 blog.yaml 선언 섹션을 순서대로 순회해 발행 Post 수집
package llms

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// collectLang gathers published posts for one language across declared sections.
func collectLang(root string, b *blogyaml.Blog, lang string) []Post {
	var posts []Post
	for _, section := range b.Sections {
		posts = append(posts, collectSection(root, lang, section)...)
	}
	return posts
}
