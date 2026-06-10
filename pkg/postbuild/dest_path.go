//ff:func feature=postbuild type=generator control=sequence
//ff:what 글 소스 경로를 빌드 산출물 옆 .md 경로로 변환 — {slug}.md → public/.../{slug}.md, 번들 index.md → 번들명.md
package postbuild

import "path/filepath"

// DestPath maps a content source file to its served .md path under publicDir,
// mirroring the page URL: /{lang}/{section}/{slug}/ -> /{lang}/{section}/{slug}.md.
func DestPath(contentDir, publicDir, src string) string {
	rel, err := filepath.Rel(contentDir, src)
	if err != nil {
		rel = src
	}
	if filepath.Base(rel) == "index.md" {
		rel = filepath.Dir(rel) + ".md"
	}
	return filepath.Join(publicDir, rel)
}
