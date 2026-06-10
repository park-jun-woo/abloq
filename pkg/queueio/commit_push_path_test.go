//ff:func feature=queueio type=client control=sequence
//ff:what commitPushPath가 변경분을 지정 메시지로 커밋·푸시(true)하고 깨끗한 트리는 no-op(false), 다른 경로 변경은 무시하는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommitPushPath(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	// Clean tree — the idempotent no-op.
	committed, err := commitPushPath(cfg, "reports/x.md", "msg")
	if err != nil || committed {
		t.Fatalf("clean tree must be a no-op: %v %v", committed, err)
	}
	if err := os.MkdirAll(filepath.Join(cfg.Workdir, "reports"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cfg.Workdir, "reports", "x.md"), []byte("x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// A change outside relPath stays uncommitted.
	if err := os.WriteFile(filepath.Join(cfg.Workdir, "other.txt"), []byte("y\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	committed, err = commitPushPath(cfg, "reports/x.md", "publish x")
	if err != nil || !committed {
		t.Fatalf("change under relPath must commit: %v %v", committed, err)
	}
	if msg := mustGit(t, cfg.Workdir, "log", "-1", "--format=%s"); msg != "publish x" {
		t.Errorf("commit message: want %q, got %q", "publish x", msg)
	}
	if status := mustGit(t, cfg.Workdir, "status", "--porcelain", "--", "other.txt"); status == "" {
		t.Error("a change outside relPath must stay uncommitted")
	}
}
