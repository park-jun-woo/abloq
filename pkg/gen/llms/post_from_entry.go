//ff:func feature=gen type=parser control=sequence
//ff:what 콘텐츠 디렉토리 엔트리 1개(글.md 또는 번들/index.md)를 발행 Post로 변환 — draft·_index·비마크다운은 false
package llms

import (
	"os"
	"path/filepath"
	"strings"
)

// postFromEntry resolves one directory entry into a published Post.
// Page bundles read <name>/index.md; plain files must end in .md;
// underscore-prefixed entries (e.g. _index.md) and drafts are skipped.
func postFromEntry(sectionDir, lang, section string, entry os.DirEntry) (Post, bool) {
	name := entry.Name()
	if strings.HasPrefix(name, "_") {
		return Post{}, false
	}
	path := filepath.Join(sectionDir, name, "index.md")
	slug := name
	if !entry.IsDir() {
		if !strings.HasSuffix(name, ".md") {
			return Post{}, false
		}
		path = filepath.Join(sectionDir, name)
		slug = strings.TrimSuffix(name, ".md")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Post{}, false
	}
	fm, ok := parseFrontMatter(data)
	if !ok || fm.Draft {
		return Post{}, false
	}
	title := fm.Title
	if title == "" {
		title = slug
	}
	desc := fm.Description
	if desc == "" {
		desc = fm.Summary
	}
	return Post{Lang: lang, Section: section, Slug: slug, Title: title, Date: fm.Date, Description: desc}, true
}
