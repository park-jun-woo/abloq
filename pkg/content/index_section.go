//ff:func feature=content type=parser control=iteration dimension=1
//ff:what content/{언어}/{섹션}/ 디렉토리 엔트리를 순회해 인덱스 항목 수집 — 디렉토리 없으면 빈 목록
package content

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// indexSection gathers index entries under root/content/<lang>/<section>/.
func indexSection(root string, b *blogyaml.Blog, lang, section string) []Entry {
	sectionDir := filepath.Join(root, "content", lang, section)
	dirEntries, err := os.ReadDir(sectionDir)
	if err != nil {
		return nil
	}
	var entries []Entry
	for _, de := range dirEntries {
		if e, ok := entryFromEntry(sectionDir, b, lang, section, de); ok {
			entries = append(entries, e)
		}
	}
	return entries
}
