//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what gitChangedSet이 변경 파일만 수집하고 비 git 디렉토리에 ok=false를 반환하는지 검증
package gate

import (
	"testing"
)

func TestGitChangedSet(t *testing.T) {
	dir := t.TempDir()
	writeRepoFile(t, dir, "a.md", "one\n")
	writeRepoFile(t, dir, "b.md", "two\n")
	initGitRepo(t, dir)
	changed, ok := gitChangedSet(dir)
	if !ok || len(changed) != 0 {
		t.Fatalf("clean tree: want ok and empty set, got %v %v", changed, ok)
	}
	writeRepoFile(t, dir, "a.md", "one modified\n")
	changed, ok = gitChangedSet(dir)
	if !ok || !changed["a.md"] || changed["b.md"] {
		t.Errorf("want only a.md changed, got %v", changed)
	}
	if _, ok := gitChangedSet(t.TempDir()); ok {
		t.Error("non-git dir: want ok=false")
	}
}
