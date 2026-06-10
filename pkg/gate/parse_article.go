//ff:func feature=gate type=parser control=sequence
//ff:what 공개 ParseArticle API — blog.yaml 기준으로 글 원문 1편을 Doc으로 파싱 (퀘스트가 제출물 비교에 사용)
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// ParseArticle parses one article's content into a Doc using the blog's
// structure declaration. Quests (e.g. translation) use it to compare submissions.
func ParseArticle(b *blogyaml.Blog, lang, content string) *Doc {
	return parseDoc(buildHeadingIndex(b), lang, content)
}
