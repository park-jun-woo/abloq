//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what articleFromEntry가 index.md 없는 번들 디렉토리를 거르는지 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestArticleFromEntryNoIndex(t *testing.T) {
	hi := buildHeadingIndex(loadGateBlog(t))
	dir := t.TempDir()
	empty := filepath.Join(dir, "content", "en", "tech", "empty-bundle")
	if err := os.MkdirAll(empty, 0o755); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(filepath.Dir(empty))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if _, ok := articleFromEntry(dir, hi, "en", "tech", e); ok {
			t.Errorf("bundle without index.md must be skipped")
		}
	}
}
