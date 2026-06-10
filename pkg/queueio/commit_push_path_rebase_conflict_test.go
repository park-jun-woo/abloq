//ff:func feature=queueio type=client control=sequence
//ff:what commitPushPath가 같은 파일을 다르게 고친 경쟁 커밋의 rebase 충돌을 에러로 표면화하는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommitPushPathRebaseConflict(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	// The competing commit touches the same file with different content —
	// pull --rebase cannot resolve it and the cycle surfaces the error.
	other := filepath.Join(filepath.Dir(cfg.Workdir), "competing")
	mustGit(t, "", "clone", cfg.RepoURL, other)
	if err := os.WriteFile(filepath.Join(other, "same.md"), []byte("theirs\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, other, "add", ".")
	mustGit(t, other, "-c", "user.name=o", "-c", "user.email=o@t", "commit", "-m", "theirs")
	mustGit(t, other, "push", "origin", "HEAD")
	if err := os.WriteFile(filepath.Join(cfg.Workdir, "same.md"), []byte("mine\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := commitPushPath(cfg, "same.md", "mine"); err == nil {
		t.Error("a rebase conflict must surface as an error")
	}
}
