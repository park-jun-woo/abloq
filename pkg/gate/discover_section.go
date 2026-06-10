//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what content/{언어}/{섹션}/ 디렉토리 엔트리를 순회해 대상 글 목록 수집 — 디렉토리 없으면 빈 목록
package gate

import (
	"os"
	"path/filepath"
)

// discoverSection collects articles under dir/content/<lang>/<section>/.
func discoverSection(dir string, hi headingIndex, lang, section string) []*Article {
	entries, err := os.ReadDir(filepath.Join(dir, "content", lang, section))
	if err != nil {
		return nil
	}
	var arts []*Article
	for _, entry := range entries {
		if a, ok := articleFromEntry(dir, hi, lang, section, entry); ok {
			arts = append(arts, a)
		}
	}
	return arts
}
