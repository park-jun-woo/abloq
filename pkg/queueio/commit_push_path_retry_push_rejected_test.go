//ff:func feature=queueio type=client control=sequence
//ff:what commitPushPath가 pre-receive 훅이 모든 푸시를 거부할 때 재시도 푸시 실패를 에러로 내는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommitPushPathRetryPushRejected(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	// A pre-receive hook rejects every push: the first push fails, the
	// rebase is a clean no-op, and the retry push fails too.
	bare := strings.TrimPrefix(cfg.RepoURL, "file://")
	hook := filepath.Join(bare, "hooks", "pre-receive")
	if err := os.WriteFile(hook, []byte("#!/bin/sh\nexit 1\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cfg.Workdir, "x.md"), []byte("x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	committed, err := commitPushPath(cfg, "x.md", "m")
	if err == nil || committed {
		t.Errorf("a rejected retry push must error: %v %v", committed, err)
	}
}
