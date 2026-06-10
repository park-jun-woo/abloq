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
		"languages: [ko]\nsections: [tech]\ngeo:\n  freshness_days: 1\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("BLOG_REPO_PATH", dir)
	posts := `[{"lang":"ko","section":"tech","slug":"post-a","date":"2026-06-01","lastmod":"2026-06-05"}]`
	resp, err := Freshness(FreshnessRequest{PostsJSON: posts, HitsJSON: "[]"})
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
	// Crawl hits flow into the priority.
	hits := `[{"lang":"ko","section":"tech","slug":"post-a","hits":42}]`
	resp, err = Freshness(FreshnessRequest{PostsJSON: posts, HitsJSON: hits})
	if err != nil {
		t.Fatal(err)
	}
	rows, _ = pqueueio.DecodeRows(resp.ItemsJSON)
	if rows[0].Priority != 42 {
		t.Errorf("hits sum must win the priority: %d", rows[0].Priority)
	}
	t.Setenv("BLOG_REPO_PATH", "")
	if _, err := Freshness(FreshnessRequest{PostsJSON: "[]", HitsJSON: "[]"}); err == nil {
		t.Error("missing BLOG_REPO_PATH must error")
	}
	var raw []map[string]any
	if err := json.Unmarshal(resp.ItemsJSON, &raw); err != nil {
		t.Fatal(err)
	}
	if _, hasID := raw[0]["id"]; hasID {
		t.Error("insert JSON must not carry DB ids")
	}
}
