package scan

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestFreshness(t *testing.T) {
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"  default_lang_in_subdir: false\nlanguages: [ko]\nsections: [tech]\ngeo:\n  freshness_days: 1\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	// One published article so the URL map can attribute the GSC page.
	post := "---\ntitle: A\ndate: 2026-06-01\nlastmod: 2026-06-05\ndraft: false\n---\n\nbody\n"
	if err := os.MkdirAll(filepath.Join(dir, "content", "ko", "tech"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "content", "ko", "tech", "post-a.md"), []byte(post), 0o644); err != nil {
		t.Fatal(err)
	}
	posts := `[{"lang":"ko","section":"tech","slug":"post-a","date":"2026-06-01","lastmod":"2026-06-05"}]`
	empty := FreshnessRequest{RepoPath: dir, PostsJSON: posts, HitsJSON: "[]", BotsJSON: "[]", GscJSON: "[]", CitesJSON: "[]"}
	resp, err := Freshness(empty)
	if err != nil {
		t.Fatalf("Freshness: %v", err)
	}
	if resp.Detected != 1 {
		t.Fatalf("want 1 detected, got %d", resp.Detected)
	}
	rows, err := pqueueio.DecodeRows(resp.ItemsJSON)
	if err != nil || len(rows) != 1 {
		t.Fatalf("ItemsJSON must decode to 1 row: %v", err)
	}
	if rows[0].Kind != "refresh" || rows[0].Section != "tech" {
		t.Errorf("unexpected row: %+v", rows[0])
	}
	// No measurement signals at all — the cold-start date score (epoch days).
	if rows[0].Priority != 20605 {
		t.Errorf("cold-start fallback must keep the date score untouched: %d", rows[0].Priority)
	}
	// All-time crawl hits flow into the cold-start priority.
	req := empty
	req.HitsJSON = `[{"lang":"ko","section":"tech","slug":"post-a","hits":42}]`
	resp, err = Freshness(req)
	if err != nil {
		t.Fatal(err)
	}
	rows, _ = pqueueio.DecodeRows(resp.ItemsJSON)
	if rows[0].Priority != 42 {
		t.Errorf("hits sum must win the cold-start priority: %d", rows[0].Priority)
	}
	// Measured window signals route to the weighted score (defaults
	// fetcher=3, train=1, gsc=1, citation=2): 3*3 + 1*7 + 1*120 + 2*2 = 140.
	req.BotsJSON = `[{"bot":"GPTBot","lang":"ko","section":"tech","slug":"post-a","hits":7,"md_hits":0},` +
		`{"bot":"ChatGPT-User","lang":"ko","section":"tech","slug":"post-a","hits":3,"md_hits":0}]`
	req.GscJSON = `[{"page":"https://t.example.com/tech/post-a/","impressions":120,"clicks":8}]`
	req.CitesJSON = `[{"lang":"ko","section":"tech","slug":"post-a","cited":2,"total":3}]`
	resp, err = Freshness(req)
	if err != nil {
		t.Fatal(err)
	}
	rows, _ = pqueueio.DecodeRows(resp.ItemsJSON)
	if rows[0].Priority != 140 {
		t.Errorf("measured priority must be the weighted sum 140: %d", rows[0].Priority)
	}
	noRepo := empty
	noRepo.RepoPath = ""
	if _, err := Freshness(noRepo); err == nil {
		t.Error("missing repo_path must error")
	}
	var raw []map[string]any
	if err := json.Unmarshal(resp.ItemsJSON, &raw); err != nil {
		t.Fatal(err)
	}
	if _, hasID := raw[0]["id"]; hasID {
		t.Error("insert JSON must not carry DB ids")
	}
}
