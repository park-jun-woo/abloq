//ff:func feature=scan type=parser control=sequence topic=cluster
//ff:what 대상 글 1편 → 그래프 노드 — front matter 부분집합 디코드(slug 오버라이드·draft 제외), 아웃링크는 미해석 대상 그대로
package cluster

import (
	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/gate"
)

// postFromArticle resolves one discovered article into a graph node. Drafts
// are out of scope (the posts index also skips them); broken front matter
// stays in — its cluster decay is still the corpus' decay. Outlinks holds the
// resolved targets before corpus filtering (Scan finalizes the edge set once
// every node is known).
func postFromArticle(b *blogyaml.Blog, lang string, a *gate.Article) (post, bool) {
	var fm frontMatter
	_ = yaml.Unmarshal([]byte(a.Doc.FrontMatter), &fm)
	if fm.Draft {
		return post{}, false
	}
	slug := a.Slug
	if fm.Slug != "" {
		slug = fm.Slug
	}
	tags := fm.Tags
	if tags == nil {
		tags = []string{}
	}
	return post{
		Section:  a.Section,
		Slug:     slug,
		Date:     fm.Date,
		Tags:     tags,
		Outlinks: linkTargets(b, lang, a.Doc.Body),
	}, true
}
