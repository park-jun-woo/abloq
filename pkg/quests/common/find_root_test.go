//ff:func feature=quest type=parser control=sequence
//ff:what FindRoot 검증 — 하위 디렉토리에서 blog.yaml 보유 조상으로 상승, 미보유 트리는 에러
package common

import (
	"path/filepath"
	"testing"
)

func TestFindRoot(t *testing.T) {
	root, abs := writeFixture(t, "content/en/posts/a.md", "x\n")
	got, err := FindRoot(filepath.Dir(abs))
	if err != nil {
		t.Fatalf("FindRoot: %v", err)
	}
	if got != root {
		t.Errorf("root = %q, want %q", got, root)
	}
	if _, err := FindRoot(t.TempDir()); err == nil {
		t.Error("rootless tree: want error")
	}
}
