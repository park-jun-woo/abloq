//ff:func feature=queueio type=client control=sequence
//ff:what Export 1회전 — open 파일 푸시·exported id 반환, 에이전트가 파일을 지운 뒤 다음 회전에서 consumed id 검출을 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExport(t *testing.T) {
	cfg := bareFixture(t)
	open := []Row{
		{ID: 1, Item: Item{Kind: "refresh", Slug: "post-a", Lang: "ko", Section: "tech", Priority: 20605}},
		{ID: 2, Item: Item{Kind: "refresh", Slug: "post-b", Lang: "ko", Section: "tech", Priority: 20607}},
	}
	res, err := Export(cfg, open, nil)
	if err != nil {
		t.Fatalf("Export: %v", err)
	}
	if len(res.ExportedIDs) != 2 || len(res.ConsumedIDs) != 0 {
		t.Fatalf("unexpected result: %+v", res)
	}
	// The push reached the bare origin: a verification clone sees both files.
	check := filepath.Join(filepath.Dir(cfg.Workdir), "check")
	mustGit(t, "", "clone", cfg.RepoURL, check)
	if _, err := os.Stat(filepath.Join(check, "quests", "queue", Filename(open[0].Item))); err != nil {
		t.Fatalf("exported file missing in origin: %v", err)
	}
	// Agent consumes post-b: deletes the queue file and pushes.
	mustGit(t, check, "rm", filepath.Join("quests", "queue", Filename(open[1].Item)))
	mustGit(t, check, "-c", "user.name=a", "-c", "user.email=a@test", "commit", "-m", "consume post-b")
	mustGit(t, check, "push", "origin", "HEAD")
	// Next cycle: rows are now exported; post-b's deletion syncs to consumed.
	exported := []Row{
		{ID: 1, Item: open[0].Item},
		{ID: 2, Item: open[1].Item},
	}
	res, err = Export(cfg, nil, exported)
	if err != nil {
		t.Fatalf("second Export: %v", err)
	}
	if len(res.ConsumedIDs) != 1 || res.ConsumedIDs[0] != 2 {
		t.Errorf("want consumed [2], got %v", res.ConsumedIDs)
	}
	if len(res.ExportedIDs) != 0 {
		t.Errorf("no open rows: want no exported ids, got %v", res.ExportedIDs)
	}
	// A broken origin fails the cycle before any transition id is reported.
	bad := cfg
	bad.RepoURL = "file://" + filepath.Join(filepath.Dir(cfg.Workdir), "no-such-origin.git")
	bad.Workdir = filepath.Join(filepath.Dir(cfg.Workdir), "bad-work")
	if _, err := Export(bad, open, nil); err == nil {
		t.Error("missing origin must error")
	}
}
