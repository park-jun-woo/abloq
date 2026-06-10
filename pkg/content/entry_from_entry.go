//ff:func feature=content type=parser control=sequence
//ff:what 콘텐츠 디렉토리 엔트리 1개(글.md 또는 번들/index.md)를 인덱스 항목으로 변환 — draft·_index·비마크다운은 false
package content

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// entryFromEntry resolves one directory entry into an index Entry. Page
// bundles read <name>/index.md; plain files must end in .md;
// underscore-prefixed entries (e.g. _index.md) and drafts are skipped.
// A front matter slug overrides the file stem; Lastmod falls back to Date.
func entryFromEntry(sectionDir string, b *blogyaml.Blog, lang, section string, entry os.DirEntry) (Entry, bool) {
	name := entry.Name()
	if strings.HasPrefix(name, "_") {
		return Entry{}, false
	}
	path := filepath.Join(sectionDir, name, "index.md")
	slug := name
	if !entry.IsDir() {
		if !strings.HasSuffix(name, ".md") {
			return Entry{}, false
		}
		path = filepath.Join(sectionDir, name)
		slug = strings.TrimSuffix(name, ".md")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Entry{}, false
	}
	fm, body, ok := parseFrontMatter(data)
	if !ok || fm.Draft {
		return Entry{}, false
	}
	if fm.Slug != "" {
		slug = fm.Slug
	}
	title := fm.Title
	if title == "" {
		title = slug
	}
	lastmod := fm.Lastmod
	if lastmod == "" {
		lastmod = fm.Date
	}
	tags := fm.Tags
	if tags == nil {
		tags = []string{}
	}
	return Entry{
		Lang:          lang,
		Section:       section,
		Slug:          slug,
		Title:         title,
		Date:          fm.Date,
		Lastmod:       lastmod,
		WordCount:     wordCount(body),
		Tags:          tags,
		InternalLinks: internalLinks(body, b.Site.BaseURL),
		SourceCount:   sourceCount(body, b.Structure.Headings["sources"][lang]),
		URL:           entryURL(b, lang, section, slug),
	}, true
}
