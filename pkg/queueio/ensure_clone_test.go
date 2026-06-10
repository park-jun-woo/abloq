//ff:func feature=queueio type=client control=sequence
//ff:what ensureClone이 최초엔 clone, 이후엔 pull --rebase로 원격 커밋을 따라잡는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureClone(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatalf("initial clone: %v", err)
	}
	if _, err := os.Stat(filepath.Join(cfg.Workdir, "README.md")); err != nil {
		t.Fatalf("clone missing seed file: %v", err)
	}
	// Push a second commit from another clone, then ensure pull picks it up.
	other := filepath.Join(filepath.Dir(cfg.Workdir), "other")
	mustGit(t, "", "clone", cfg.RepoURL, other)
	if err := os.WriteFile(filepath.Join(other, "agent.txt"), []byte("x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, other, "add", ".")
	mustGit(t, other, "-c", "user.name=a", "-c", "user.email=a@test", "commit", "-m", "agent")
	mustGit(t, other, "push", "origin", "HEAD")
	if err := ensureClone(cfg); err != nil {
		t.Fatalf("pull --rebase: %v", err)
	}
	if _, err := os.Stat(filepath.Join(cfg.Workdir, "agent.txt")); err != nil {
		t.Error("existing clone must fast-forward to the agent commit")
	}
}
