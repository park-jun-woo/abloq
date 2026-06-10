//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 신선도 큐(quests/queue/)에 글 키가 등재됐는지 검사 — 큐 디렉토리 미도입 시 통과(스킵)
package gate

import (
	"bytes"
	"os"
	"path/filepath"
)

// queueAllows reports whether the freshness queue permits a lastmod update for
// the article key (<lang>/<section>/<slug>). A blog without quests/queue/ has
// not adopted the queue yet, so the check is skipped (allowed).
func queueAllows(dir, key string) bool {
	queueDir := filepath.Join(dir, "quests", "queue")
	entries, err := os.ReadDir(queueDir)
	if err != nil {
		return true
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(queueDir, e.Name()))
		if err == nil && bytes.Contains(data, []byte(key)) {
			return true
		}
	}
	return false
}
