//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 신선도 큐(quests/queue/)에 글 키가 등재됐는지 검사 — strconv.Quote 정확 매칭, 큐 디렉토리 미도입 시 통과(스킵)
//ff:why Phase018: 부분문자열 매칭은 en/tech/a ⊂ en/tech/abc 과허용 — 직렬화가 키를 항상 따옴표로 감싸므로(key:·keys:) 인용된 키 전체를 찾으면 정확 매칭이 된다
package gate

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
)

// queueAllows reports whether the freshness queue permits a lastmod update for
// the article key (<lang>/<section>/<slug>). The key is matched in its quoted
// form (strconv.Quote — the queue serialization's own framing on key: and
// keys: lines), so a key can never match inside a longer key. A blog without
// quests/queue/ has not adopted the queue yet, so the check is skipped
// (allowed).
func queueAllows(dir, key string) bool {
	queueDir := filepath.Join(dir, "quests", "queue")
	entries, err := os.ReadDir(queueDir)
	if err != nil {
		return true
	}
	quoted := []byte(strconv.Quote(key))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(queueDir, e.Name()))
		if err == nil && bytes.Contains(data, quoted) {
			return true
		}
	}
	return false
}
