//ff:func feature=scan type=parser control=iteration dimension=1 topic=cluster
//ff:what 기본 언어 발행 글 전수를 그래프 노드로 수집 — 언어를 기본 언어로 좁힌 Blog 사본으로 gate.Discover 재사용, draft 제외
//ff:why 그래프는 기본 언어 기준 1회 구성(Phase011 결정) — 번역본은 기본 언어 글의 클러스터 결정을 따르므로 언어별 그래프를 만들지 않는다
package cluster

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/gate"
)

// collectPosts gathers the default-language published articles as graph
// nodes, in the deterministic gate.Discover order (declared sections,
// directory-name order). Discovery runs on a Blog copy narrowed to the
// default language so translation trees are never walked.
func collectPosts(root string, b *blogyaml.Blog, lang string) []post {
	narrowed := *b
	narrowed.Languages = []string{lang}
	posts := make([]post, 0)
	for _, a := range gate.Discover(root, &narrowed) {
		if p, ok := postFromArticle(b, lang, a); ok {
			posts = append(posts, p)
		}
	}
	return posts
}
