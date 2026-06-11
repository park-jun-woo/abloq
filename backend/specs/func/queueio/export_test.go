package queueio

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func mustRun(t *testing.T, dir string, name string, args ...string) {
	t.Helper()
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%s %v: %v: %s", name, args, err, out)
	}
}

func TestExport(t *testing.T) {
	root := t.TempDir()
	bare := filepath.Join(root, "origin.git")
	seed := filepath.Join(root, "seed")
	mustRun(t, "", "git", "init", "--bare", "-b", "main", bare)
	mustRun(t, "", "git", "init", "-b", "main", seed)
	if err := os.WriteFile(filepath.Join(seed, "README.md"), []byte("seed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustRun(t, seed, "git", "add", ".")
	mustRun(t, seed, "git", "-c", "user.name=s", "-c", "user.email=s@t", "commit", "-m", "seed")
	mustRun(t, seed, "git", "push", "file://"+bare, "main")

	t.Setenv("QUEUE_EXPORT_WORKDIR", filepath.Join(root, "work"))
	open := `[{"id":1,"kind":"refresh","slug":"post-a","lang":"ko",` +
		`"payload":{"section":"tech","lastmod":"2026-06-05"},"priority":20605}]`
	resp, err := Export(ExportRequest{SiteName: "default", RepoURL: "file://" + bare,
		OpenJSON: open, ExportedJSON: "[]"})
	if err != nil {
		t.Fatalf("Export: %v", err)
	}
	if resp.Exported != 1 || resp.Consumed != 0 {
		t.Errorf("want exported=1 consumed=0, got %+v", resp)
	}
	if string(resp.ExportedIdsJSON) != "[1]" || string(resp.ConsumedIdsJSON) != "[]" {
		t.Errorf("unexpected id JSON: %s / %s", resp.ExportedIdsJSON, resp.ConsumedIdsJSON)
	}
	// The pushed file reached the bare origin.
	check := filepath.Join(root, "check")
	mustRun(t, "", "git", "clone", "file://"+bare, check)
	if _, err := os.Stat(filepath.Join(check, "quests", "queue", "refresh-ko-tech-post-a.yaml")); err != nil {
		t.Fatalf("exported file missing in origin: %v", err)
	}
	// The work clone landed under <base>/<site> — per-site isolation.
	if _, err := os.Stat(filepath.Join(root, "work", "default", ".git")); err != nil {
		t.Errorf("per-site work clone missing: %v", err)
	}
	if _, err := Export(ExportRequest{SiteName: "default", OpenJSON: "[]", ExportedJSON: "[]"}); err == nil {
		t.Error("empty site repo URL must error (per-site 500 contract)")
	}
	if _, err := Export(ExportRequest{RepoURL: "file://" + bare, OpenJSON: "[]", ExportedJSON: "[]"}); err == nil {
		t.Error("missing site name must error")
	}
}

func TestWorkdirBase(t *testing.T) {
	// Env unset falls back to the image default.
	t.Setenv("QUEUE_EXPORT_WORKDIR", "")
	if got := workdirBase(); got != "/var/lib/abloqd/queue-export" {
		t.Errorf("unset env: want image default, got %q", got)
	}
	// Env set overrides (test harnesses point it at a temp dir).
	t.Setenv("QUEUE_EXPORT_WORKDIR", "/tmp/qe")
	if got := workdirBase(); got != "/tmp/qe" {
		t.Errorf("set env: want /tmp/qe, got %q", got)
	}
}
