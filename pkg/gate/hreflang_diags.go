//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 글 1편의 빌드 페이지에서 hreflang 누락 언어를 진단 — 페이지 자체가 없으면 그 사실을 진단 (기본 언어 루트 서빙 경로 인지)
package gate

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// hreflangDiags checks one built article page for hreflang completeness
// against the languages the article exists in. The page path follows
// Blog.URLLang — a root-served default language has no language directory.
func hreflangDiags(dir string, b *blogyaml.Blog, a *Article, langs []string) []blogyaml.Diagnostic {
	page := filepath.Join("public", b.URLLang(a.Lang), a.Section, effSlug(a), "index.html")
	html, err := os.ReadFile(filepath.Join(dir, page))
	if err != nil {
		return []blogyaml.Diagnostic{{File: a.Path, Line: 1, Rule: "hreflang-complete",
			Message: "built page " + page + " missing — run the site build"}}
	}
	alts := parseAlternates(string(html))
	var diags []blogyaml.Diagnostic
	for _, lang := range langs {
		if alts[lang] != "" {
			continue
		}
		diags = append(diags, blogyaml.Diagnostic{File: a.Path, Line: 1, Rule: "hreflang-complete",
			Message: page + " lacks hreflang alternate for " + lang})
	}
	return diags
}
