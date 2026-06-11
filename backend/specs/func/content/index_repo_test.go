package content

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// fixtureBlog walks upward from the package directory until it finds the
// shared fixtures/blog repository — the same test source runs from
// backend/specs/func/content and from the generated copy under
// backend/arts/backend/internal/content (different depths).
func fixtureBlog(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	for i := 0; i < 8; i++ {
		candidate := filepath.Join(dir, "fixtures", "blog")
		if _, err := os.Stat(filepath.Join(candidate, "blog.yaml")); err == nil {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatal("fixtures/blog not found in any ancestor directory")
	return ""
}

func TestIndexRepo(t *testing.T) {
	fixture := fixtureBlog(t)
	resp, err := IndexRepo(IndexRepoRequest{RepoPath: fixture})
	if err != nil {
		t.Fatalf("IndexRepo: %v", err)
	}
	if resp.Count != 2 {
		t.Errorf("Count = %d, want 2 (fixture has 2 published posts)", resp.Count)
	}
	var entries []map[string]any
	if err := json.Unmarshal(resp.EntriesJSON, &entries); err != nil || len(entries) != 2 {
		t.Errorf("EntriesJSON = %s (err=%v), want 2-entry JSON array", resp.EntriesJSON, err)
	}

	if _, err := IndexRepo(IndexRepoRequest{}); err == nil {
		t.Error("IndexRepo without a repo_path must fail")
	}

	if _, err := IndexRepo(IndexRepoRequest{RepoPath: t.TempDir()}); err == nil {
		t.Error("IndexRepo on a repo without blog.yaml must fail")
	}
}
