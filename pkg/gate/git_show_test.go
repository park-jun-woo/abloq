//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what gitShow가 HEAD 스냅샷을 반환하고 HEAD에 없는 경로에 false를 반환하는지 검증
package gate

import "testing"

func TestGitShow(t *testing.T) {
	dir := t.TempDir()
	writeRepoFile(t, dir, "a.md", "committed\n")
	initGitRepo(t, dir)
	writeRepoFile(t, dir, "a.md", "working tree\n")
	out, ok := gitShow(dir, "a.md")
	if !ok || string(out) != "committed\n" {
		t.Errorf("gitShow = (%q, %v), want committed snapshot", out, ok)
	}
	if _, ok := gitShow(dir, "missing.md"); ok {
		t.Error("path absent from HEAD: want false")
	}
}
