//ff:func feature=queueio type=client control=sequence
//ff:what commitPushPath가 경쟁 커밋으로 거부된 첫 푸시를 pull --rebase 후 재푸시로 성공시키고 양쪽 커밋이 origin에 남는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommitPushPathRebaseRetry(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	// A competing commit lands on the origin after our clone — the first
	// push is rejected, pull --rebase absorbs it and the retry push wins.
	other := filepath.Join(filepath.Dir(cfg.Workdir), "competing")
	mustGit(t, "", "clone", cfg.RepoURL, other)
	if err := os.WriteFile(filepath.Join(other, "agent.txt"), []byte("x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, other, "add", ".")
	mustGit(t, other, "-c", "user.name=o", "-c", "user.email=o@t", "commit", "-m", "competing")
	mustGit(t, other, "push", "origin", "HEAD")
	if err := os.WriteFile(filepath.Join(cfg.Workdir, "mine.md"), []byte("y\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	committed, err := commitPushPath(cfg, "mine.md", "mine")
	if err != nil || !committed {
		t.Fatalf("rebase retry must succeed: %v %v", committed, err)
	}
	// Both commits are on the origin.
	check := filepath.Join(filepath.Dir(cfg.Workdir), "check2")
	mustGit(t, "", "clone", cfg.RepoURL, check)
	if _, err := os.Stat(filepath.Join(check, "mine.md")); err != nil {
		t.Error("retried push missing from origin")
	}
	if _, err := os.Stat(filepath.Join(check, "agent.txt")); err != nil {
		t.Error("competing commit lost by the rebase")
	}
}
