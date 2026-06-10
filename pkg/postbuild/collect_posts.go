//ff:func feature=postbuild type=parser control=iteration dimension=1
//ff:what content/{lang}/{section}/의 발행 글(.md)을 수집 — 단일 파일과 번들(index.md) 포함, _index.md 제외, 정렬 반환
package postbuild

import (
	"path/filepath"
	"sort"
)

// CollectPosts lists every article source under contentDir: page files at
// {lang}/{section}/{slug}.md plus bundles at {lang}/{section}/{slug}/index.md.
func CollectPosts(contentDir string) ([]string, error) {
	pages, err := filepath.Glob(filepath.Join(contentDir, "*", "*", "*.md"))
	if err != nil {
		return nil, err
	}
	bundles, err := filepath.Glob(filepath.Join(contentDir, "*", "*", "*", "index.md"))
	if err != nil {
		return nil, err
	}
	var posts []string
	for _, p := range append(pages, bundles...) {
		if filepath.Base(p) == "_index.md" {
			continue
		}
		posts = append(posts, p)
	}
	sort.Strings(posts)
	return posts, nil
}
