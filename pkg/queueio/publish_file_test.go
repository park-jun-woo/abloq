//ff:func feature=queueio type=client control=sequence
//ff:what PublishFile이 파일을 origin까지 발행하고 동일 내용 재발행은 no-op(커밋 0·false), 변경 내용은 새 커밋인지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPublishFile(t *testing.T) {
	cfg := bareFixture(t)
	committed, err := PublishFile(cfg, "reports/2026-04.md", []byte("# report\n"))
	if err != nil {
		t.Fatalf("PublishFile: %v", err)
	}
	if !committed {
		t.Error("first publish must commit")
	}
	// The file reached the origin.
	check := filepath.Join(filepath.Dir(cfg.Workdir), "check")
	mustGit(t, "", "clone", cfg.RepoURL, check)
	data, err := os.ReadFile(filepath.Join(check, "reports", "2026-04.md"))
	if err != nil || string(data) != "# report\n" {
		t.Fatalf("origin content wrong: %q, %v", data, err)
	}
	before := mustGit(t, cfg.Workdir, "rev-parse", "HEAD")
	// Identical content — idempotent no-op.
	committed, err = PublishFile(cfg, "reports/2026-04.md", []byte("# report\n"))
	if err != nil {
		t.Fatalf("idempotent PublishFile: %v", err)
	}
	if committed {
		t.Error("identical content must not commit")
	}
	if after := mustGit(t, cfg.Workdir, "rev-parse", "HEAD"); after != before {
		t.Error("identical content must not produce a new commit")
	}
	// Changed content — a new commit.
	committed, err = PublishFile(cfg, "reports/2026-04.md", []byte("# report v2\n"))
	if err != nil || !committed {
		t.Fatalf("changed content must commit: %v %v", committed, err)
	}
}
