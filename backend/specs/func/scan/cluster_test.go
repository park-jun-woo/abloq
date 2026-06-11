package scan

import (
	"os"
	"path/filepath"
	"testing"

	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestCluster(t *testing.T) {
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [tech]\ngeo:\n  min_internal_links: 2\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	postDir := filepath.Join(dir, "content", "ko", "tech")
	if err := os.MkdirAll(postDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// hub and thin link each other: both sit at 1 outlink (below min 2) and
	// 1 inlink (not isolated); no tags, no taxonomy — exactly one
	// min-internal-links violation each.
	posts := map[string]string{
		"hub.md":  "---\ntitle: Hub\ndate: 2026-01-05\n---\n\n[t](/tech/thin/)\n",
		"thin.md": "---\ntitle: Thin\ndate: 2026-01-04\n---\n\n[h](/tech/hub/)\n",
	}
	for name, body := range posts {
		if err := os.WriteFile(filepath.Join(postDir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	resp, err := Cluster(ClusterRequest{RepoPath: dir})
	if err != nil {
		t.Fatalf("Cluster: %v", err)
	}
	if resp.Detected != 2 {
		t.Fatalf("want 2 detected, got %d", resp.Detected)
	}
	rows, err := pqueueio.DecodeRows(resp.ItemsJSON)
	if err != nil || len(rows) != 2 {
		t.Fatalf("ItemsJSON must decode to 2 rows: %v", err)
	}
	if rows[0].Kind != "cluster" || rows[0].Section != "tech" || rows[0].Payload["violations"] == "" {
		t.Errorf("unexpected row: %+v", rows[0])
	}
	if _, err := Cluster(ClusterRequest{}); err == nil {
		t.Error("missing repo_path must error")
	}
	if _, err := Cluster(ClusterRequest{RepoPath: t.TempDir()}); err == nil {
		t.Error("missing blog.yaml must error")
	}
	invalid := t.TempDir()
	if err := os.WriteFile(filepath.Join(invalid, "blog.yaml"), []byte("site:\n  baseURL: not-a-url\n  title: T\n  author: A\nlanguages: [ko]\nsections: [tech]\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Cluster(ClusterRequest{RepoPath: invalid}); err == nil {
		t.Error("invalid blog.yaml must error")
	}
}
