package report

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func gitOut(t *testing.T, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func fixtureRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"  default_lang_in_subdir: false\nlanguages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	post := "---\ntitle: A\ndate: 2026-06-01\nlastmod: 2026-06-05\ndraft: false\n---\n\nbody\n"
	if err := os.MkdirAll(filepath.Join(dir, "content", "ko", "tech"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "content", "ko", "tech", "post-a.md"), []byte(post), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func bareOrigin(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	bare := filepath.Join(root, "origin.git")
	seed := filepath.Join(root, "seed")
	gitOut(t, "", "init", "--bare", "-b", "main", bare)
	gitOut(t, "", "init", "-b", "main", seed)
	if err := os.WriteFile(filepath.Join(seed, "README.md"), []byte("seed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOut(t, seed, "add", ".")
	gitOut(t, seed, "-c", "user.name=s", "-c", "user.email=s@t", "commit", "-m", "seed")
	gitOut(t, seed, "push", "file://"+bare, "main")
	return bare, filepath.Join(root, "work")
}

func TestMonthly(t *testing.T) {
	bare, work := bareOrigin(t)
	t.Setenv("QUEUE_EXPORT_WORKDIR", work)
	req := MonthlyRequest{
		RepoPath:  fixtureRepo(t),
		SiteName:  "default",
		RepoURL:   "file://" + bare,
		Ym:        "2026-04",
		PostsJSON: `[{"lang":"ko","section":"tech","slug":"post-a","date":"2026-06-01","lastmod":"2026-06-05"}]`,
		BotsJSON: `[{"bot":"GPTBot","lang":"ko","section":"tech","slug":"post-a","hits":7,"md_hits":2},` +
			`{"bot":"ChatGPT-User","lang":"ko","section":"tech","slug":"post-a","hits":3,"md_hits":0}]`,
		PrevBotsJSON:  `[{"bot":"GPTBot","lang":"ko","section":"tech","slug":"post-a","hits":5,"md_hits":0}]`,
		GscJSON:       `[{"page":"https://t.example.com/tech/post-a/","impressions":120,"clicks":8}]`,
		PrevGscJSON:   "[]",
		CitesJSON:     `[{"lang":"ko","section":"tech","slug":"post-a","cited":2,"total":3}]`,
		PrevCitesJSON: "[]",
		QueueJSON:     "[]",
		UnknownJSON:   `[{"ua":"PetalBot","hits":1}]`,
	}
	resp, err := Monthly(req)
	if err != nil {
		t.Fatalf("Monthly: %v", err)
	}
	if resp.Ym != "2026-04" || resp.Articles != 1 || !resp.Published {
		t.Fatalf("unexpected response: %+v", resp)
	}
	// Weighted priority with defaults (3/1/1/2): 3*3 + 1*7 + 1*120 + 2*2 = 140.
	if !strings.Contains(resp.Markdown, "| ko/tech/post-a | 2026-06-01 | 7 | 0 | 3 | 2 | 120 | 8 | 2 | 140 |") {
		t.Errorf("markdown row wrong:\n%s", resp.Markdown)
	}
	if !strings.Contains(resp.ReportJSON, `"priority":140`) {
		t.Errorf("report JSON misses the priority: %s", resp.ReportJSON)
	}
	// The publication copy reached the origin under reports/<ym>.md.
	check := t.TempDir()
	gitOut(t, "", "clone", "file://"+bare, filepath.Join(check, "c"))
	data, err := os.ReadFile(filepath.Join(check, "c", "reports", "2026-04.md"))
	if err != nil || !strings.Contains(string(data), "# Visibility report 2026-04") {
		t.Fatalf("publication copy wrong: %v", err)
	}
	// Regeneration is the idempotent no-op (identical markdown, no commit).
	resp2, err := Monthly(req)
	if err != nil {
		t.Fatalf("regenerate: %v", err)
	}
	if resp2.Published {
		t.Error("identical regeneration must not publish a new commit")
	}
	if resp2.Markdown != resp.Markdown || resp2.ReportJSON != resp.ReportJSON {
		t.Error("regeneration must be byte-identical")
	}
	// Default ym ('') resolves to the last closed month and still works.
	req.Ym = ""
	resp3, err := Monthly(req)
	if err != nil {
		t.Fatalf("default ym: %v", err)
	}
	if resp3.Ym == "" || len(resp3.Ym) != 7 {
		t.Errorf("default ym must resolve to YYYY-MM: %q", resp3.Ym)
	}
	// Malformed ym is rejected.
	req.Ym = "2026-4"
	if _, err := Monthly(req); err == nil {
		t.Error("malformed ym must error")
	}
}

func TestMonthlyErrors(t *testing.T) {
	t.Setenv("QUEUE_EXPORT_WORKDIR", filepath.Join(t.TempDir(), "w"))
	valid := MonthlyRequest{SiteName: "default", Ym: "2026-04",
		PostsJSON: "[]", BotsJSON: "[]", PrevBotsJSON: "[]",
		GscJSON: "[]", PrevGscJSON: "[]", CitesJSON: "[]", PrevCitesJSON: "[]",
		QueueJSON: "[]", UnknownJSON: "[]"}
	if _, err := Monthly(valid); err == nil {
		t.Error("missing repo_path must error")
	}
	// Invalid blog.yaml (negative weight) is rejected before any work.
	bad := t.TempDir()
	if err := os.WriteFile(filepath.Join(bad, "blog.yaml"), []byte(
		"site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n"+
			"languages: [ko]\nsections: [tech]\ngeo:\n  priority_weights:\n    fetcher: -1\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	broken := valid
	broken.RepoPath = bad
	if _, err := Monthly(broken); err == nil {
		t.Error("invalid blog.yaml must error")
	}
	valid.RepoPath = fixtureRepo(t)
	// Malformed aggregate JSON surfaces as a decode error.
	broken = valid
	broken.BotsJSON = "not json"
	if _, err := Monthly(broken); err == nil {
		t.Error("malformed aggregate JSON must error")
	}
	// The site row's publication repo URL is required (publish-on-generate).
	if _, err := Monthly(valid); err == nil {
		t.Error("missing queue_export_repo_url must error")
	}
	// A missing site name cannot derive the per-site work clone.
	broken = valid
	broken.SiteName = ""
	broken.RepoURL = "file:///nonexistent/origin.git"
	if _, err := Monthly(broken); err == nil {
		t.Error("missing site name must error")
	}
	// An unreachable origin fails the publication step.
	valid.RepoURL = "file:///nonexistent/origin.git"
	if _, err := Monthly(valid); err == nil {
		t.Error("unreachable origin must error")
	}
}

func TestWorkdirBase(t *testing.T) {
	// Env unset falls back to the image default — same convention as the
	// queue exporter.
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
