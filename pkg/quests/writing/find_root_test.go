//ff:func feature=quest type=parser control=sequence
//ff:what findRoot 검증 — 하위 디렉토리에서 blog.yaml 보유 조상을 찾고, 없으면 에러
package writing

import (
	"path/filepath"
	"testing"
)

func TestFindRoot(t *testing.T) {
	root := writeInstance(t)
	writeFile(t, root, "content/en/posts/.keep", "")
	got, err := findRoot(filepath.Join(root, "content", "en", "posts"))
	if err != nil {
		t.Fatalf("findRoot: %v", err)
	}
	if got != root {
		t.Errorf("root = %q, want %q", got, root)
	}
	if _, err := findRoot(t.TempDir()); err == nil {
		t.Error("rootless dir: want error, got nil")
	}
}
