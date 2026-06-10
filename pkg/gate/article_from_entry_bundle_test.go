//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what articleFromEntry가 페이지 번들(<이름>/index.md)을 디렉토리명 slug로 채택하는지 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestArticleFromEntryBundle(t *testing.T) {
	hi := buildHeadingIndex(loadGateBlog(t))
	dir := t.TempDir()
	bundle := filepath.Join(dir, "content", "en", "tech", "my-post")
	if err := os.MkdirAll(bundle, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(bundle, "index.md"), []byte("---\ntitle: B\n---\nbody\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(filepath.Dir(bundle))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		a, ok := articleFromEntry(dir, hi, "en", "tech", e)
		if !ok || a.Slug != "my-post" {
			t.Errorf("bundle entry = (%v, %v), want slug my-post", a, ok)
		}
	}
}
