//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what blog.yaml 선언 언어를 순서대로 순회해 발행 글(draft:false) 전체 수집 — llms.txt 생성의 입력
package llms

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Collect gathers all published posts under root/content/ for the declared
// languages and sections. Undeclared directories are ignored.
func Collect(root string, b *blogyaml.Blog) []Post {
	var posts []Post
	for _, lang := range b.Languages {
		posts = append(posts, collectLang(root, b, lang)...)
	}
	return posts
}
