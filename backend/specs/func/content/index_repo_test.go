//ff:func feature=content type=parser control=sequence
//ff:what IndexRepo 래퍼가 BLOG_REPO_PATH의 픽스처 저장소를 인덱싱해 Count·EntriesJSON을 채우고, 미설정이면 에러인지 검증
package content

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestIndexRepo(t *testing.T) {
	fixture, err := filepath.Abs(filepath.Join("..", "..", "..", "fixtures", "blog"))
	if err != nil {
		t.Fatalf("fixture path: %v", err)
	}
	t.Setenv("BLOG_REPO_PATH", fixture)
	resp, err := IndexRepo(IndexRepoRequest{})
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

	t.Setenv("BLOG_REPO_PATH", "")
	if _, err := IndexRepo(IndexRepoRequest{}); err == nil {
		t.Error("IndexRepo without BLOG_REPO_PATH must fail")
	}

	t.Setenv("BLOG_REPO_PATH", t.TempDir())
	if _, err := IndexRepo(IndexRepoRequest{}); err == nil {
		t.Error("IndexRepo on a repo without blog.yaml must fail")
	}
}
