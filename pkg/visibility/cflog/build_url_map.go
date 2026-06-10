//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 저장소 인덱스에서 URI 역매핑 구성 — 글마다 페이지 경로(말미 슬래시·index.html)와 .md 병행 서빙 경로를 등록
//ff:why URI→글 매핑은 저장소가 단일 소스다(posts @get 안 씀 — Phase010/011 선례): pkg/content 인덱스와 Blog.URLLang·postbuild .md 규칙에서 역으로 푼다. 매핑 안 되는 URI(정적 자산·404 등)는 수집기가 버린다 (Phase012)
package cflog

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/content"
)

// BuildURLMap indexes the blog repository at root and registers, per
// article, the served request paths: /{lang?}/{section}/{slug}/ (plus its
// index.html form) and the parallel-served /{lang?}/{section}/{slug}.md.
// The language segment follows Blog.URLLang — a root-served default
// language drops it.
func BuildURLMap(root string, b *blogyaml.Blog) (map[string]Article, error) {
	entries, err := content.IndexRepo(root)
	if err != nil {
		return nil, err
	}
	urls := make(map[string]Article, 3*len(entries))
	for _, e := range entries {
		base := "/"
		if seg := b.URLLang(e.Lang); seg != "" {
			base += seg + "/"
		}
		base += e.Section + "/" + e.Slug
		page := Article{Lang: e.Lang, Section: e.Section, Slug: e.Slug}
		md := Article{Lang: e.Lang, Section: e.Section, Slug: e.Slug, MD: true}
		urls[base+"/"] = page
		urls[base+"/index.html"] = page
		urls[base+".md"] = md
	}
	return urls, nil
}
