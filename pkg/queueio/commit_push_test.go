//ff:func feature=queueio type=client control=sequence
//ff:what commitPush가 변경분을 author env 정체성으로 푸시하고 동일 내용 재실행은 no-op(커밋 0)인지 검증
package queueio

import (
	"path/filepath"
	"testing"
)

func TestCommitPush(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	it := Item{Kind: "refresh", Slug: "post-a", Lang: "ko", Section: "tech"}
	if err := WriteDir(filepath.Join(cfg.Workdir, "quests", "queue"), []Item{it}); err != nil {
		t.Fatal(err)
	}
	if err := commitPush(cfg); err != nil {
		t.Fatalf("commitPush: %v", err)
	}
	if author := mustGit(t, cfg.Workdir, "log", "-1", "--format=%an"); author != "abloqd-bot" {
		t.Errorf("want author abloqd-bot, got %s", author)
	}
	before := mustGit(t, cfg.Workdir, "rev-parse", "HEAD")
	// Same content again — must be a no-op (no new commit, no error).
	if err := WriteDir(filepath.Join(cfg.Workdir, "quests", "queue"), []Item{it}); err != nil {
		t.Fatal(err)
	}
	if err := commitPush(cfg); err != nil {
		t.Fatalf("idempotent commitPush: %v", err)
	}
	if after := mustGit(t, cfg.Workdir, "rev-parse", "HEAD"); after != before {
		t.Error("identical content must not produce a new commit")
	}
}
