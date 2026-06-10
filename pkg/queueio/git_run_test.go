//ff:func feature=queueio type=client control=sequence
//ff:what gitRun이 stdout을 트림해 반환하고 실패 시 git stderr를 에러 메시지에 싣는지 검증
package queueio

import (
	"strings"
	"testing"
)

func TestGitRun(t *testing.T) {
	dir := t.TempDir()
	if _, err := gitRun(dir, "init", "-b", "main"); err != nil {
		t.Fatalf("git init: %v", err)
	}
	out, err := gitRun(dir, "rev-parse", "--is-inside-work-tree")
	if err != nil || out != "true" {
		t.Errorf("want true, got %q (%v)", out, err)
	}
	if _, err := gitRun(dir, "rev-parse", "definitely-no-such-ref"); err == nil {
		t.Error("bad ref must error")
	} else if !strings.Contains(err.Error(), "git rev-parse") {
		t.Errorf("error must name the command: %v", err)
	}
}
