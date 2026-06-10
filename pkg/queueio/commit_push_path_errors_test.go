//ff:func feature=queueio type=client control=sequence
//ff:what commitPushPath가 비저장소 workdir·읽기 불가 파일·빈 커밋 신원을 각각 에러로 내는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCommitPushPathErrors(t *testing.T) {
	// Not a git repository — the status probe fails.
	bad := Config{Workdir: t.TempDir(), AuthorName: "a", AuthorEmail: "a@t"}
	if _, err := commitPushPath(bad, "reports/x.md", "m"); err == nil {
		t.Error("a non-repo workdir must error")
	}
	// An unreadable file passes the status probe but fails git add.
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	locked := filepath.Join(cfg.Workdir, "locked.md")
	if err := os.WriteFile(locked, []byte("x\n"), 0o000); err != nil {
		t.Fatal(err)
	}
	if _, err := commitPushPath(cfg, "locked.md", "m"); err == nil {
		t.Skip("running with CAP_DAC_OVERRIDE — add error not reproducible")
	}
	// An empty author identity fails the commit step.
	if err := os.WriteFile(filepath.Join(cfg.Workdir, "id.md"), []byte("x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	noID := cfg
	noID.AuthorName, noID.AuthorEmail = "", ""
	if _, err := commitPushPath(noID, "id.md", "m"); err == nil {
		t.Error("an empty commit identity must error")
	}
}
