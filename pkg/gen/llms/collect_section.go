//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what content/{언어}/{섹션}/ 디렉토리 엔트리를 순회해 발행 Post 목록 수집 — 디렉토리 없으면 빈 목록
package llms

import (
	"os"
	"path/filepath"
)

// collectSection gathers published posts under root/content/<lang>/<section>/.
func collectSection(root, lang, section string) []Post {
	sectionDir := filepath.Join(root, "content", lang, section)
	entries, err := os.ReadDir(sectionDir)
	if err != nil {
		return nil
	}
	var posts []Post
	for _, entry := range entries {
		if p, ok := postFromEntry(sectionDir, lang, section, entry); ok {
			posts = append(posts, p)
		}
	}
	return posts
}
