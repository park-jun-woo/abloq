//ff:func feature=postbuild type=generator control=sequence
//ff:what 글 소스 경로를 빌드 산출물 옆 .md 경로로 변환 — {slug}.md → public/.../{slug}.md, 번들 index.md → 번들명.md, 루트 서빙 기본 언어는 언어 디렉토리 생략
package postbuild

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// DestPath maps a content source file to its served .md path under publicDir,
// mirroring the page URL: /{lang}/{section}/{slug}/ -> /{lang}/{section}/{slug}.md.
// The language directory follows Blog.URLLang — a root-served default language
// (site.default_lang_in_subdir: false) drops its language segment.
func DestPath(contentDir, publicDir, src string, b *blogyaml.Blog) string {
	rel, err := filepath.Rel(contentDir, src)
	if err != nil {
		rel = src
	}
	if filepath.Base(rel) == "index.md" {
		rel = filepath.Dir(rel) + ".md"
	}
	parts := strings.SplitN(filepath.ToSlash(rel), "/", 2)
	if len(parts) == 2 && b.URLLang(parts[0]) == "" {
		rel = filepath.FromSlash(parts[1])
	}
	return filepath.Join(publicDir, rel)
}
