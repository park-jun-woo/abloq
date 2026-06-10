//ff:func feature=gate type=frame control=sequence
//ff:what 콘텐츠 디렉토리 엔트리 1개(글.md 또는 번들/index.md)를 대상 글로 변환 — _index·비마크다운은 false
package gate

import (
	"os"
	"path/filepath"
	"strings"
)

// articleFromEntry resolves one directory entry into an Article. Page bundles
// read <name>/index.md; plain files must end in .md; underscore-prefixed
// entries (e.g. _index.md) are skipped.
func articleFromEntry(dir string, hi headingIndex, lang, section string, entry os.DirEntry) (*Article, bool) {
	name := entry.Name()
	if strings.HasPrefix(name, "_") {
		return nil, false
	}
	rel := filepath.Join("content", lang, section, name, "index.md")
	slug := name
	if !entry.IsDir() {
		if !strings.HasSuffix(name, ".md") {
			return nil, false
		}
		rel = filepath.Join("content", lang, section, name)
		slug = strings.TrimSuffix(name, ".md")
	}
	data, err := os.ReadFile(filepath.Join(dir, rel))
	if err != nil {
		return nil, false
	}
	doc := parseDoc(hi, lang, string(data))
	return &Article{Lang: lang, Section: section, Slug: slug, Path: rel, Doc: doc}, true
}
