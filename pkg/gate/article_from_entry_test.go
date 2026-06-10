//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what articleFromEntry가 .md만 채택하고 _index·비마크다운을 거르는지 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestArticleFromEntry(t *testing.T) {
	hi := buildHeadingIndex(loadGateBlog(t))
	dir := t.TempDir()
	secDir := filepath.Join(dir, "content", "en", "tech")
	if err := os.MkdirAll(secDir, 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{"a.md": "---\ntitle: A\n---\nbody\n", "_index.md": "x", "img.png": "x"}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(secDir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	entries, err := os.ReadDir(secDir)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, e := range entries {
		if a, ok := articleFromEntry(dir, hi, "en", "tech", e); ok {
			got = append(got, a.Slug)
		}
	}
	if len(got) != 1 || got[0] != "a" {
		t.Errorf("accepted slugs = %v, want [a]", got)
	}
}
