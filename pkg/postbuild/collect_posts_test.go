//ff:func feature=postbuild type=parser control=iteration dimension=1
//ff:what CollectPosts가 단일 파일과 번들 index.md를 모으고 _index.md를 제외하며 정렬하는지 검증
package postbuild

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCollectPosts(t *testing.T) {
	dir := t.TempDir()
	files := []string{
		"ko/tech/post.md",
		"ko/tech/_index.md",
		"ko/tech/bundle/index.md",
		"en/opinion/essay.md",
	}
	for _, f := range files {
		p := filepath.Join(dir, f)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatalf("MkdirAll: %v", err)
		}
		if err := os.WriteFile(p, []byte("x"), 0o644); err != nil {
			t.Fatalf("write %s: %v", f, err)
		}
	}
	got, err := CollectPosts(dir)
	if err != nil {
		t.Fatalf("CollectPosts: %v", err)
	}
	want := []string{
		filepath.Join(dir, "en/opinion/essay.md"),
		filepath.Join(dir, "ko/tech/bundle/index.md"),
		filepath.Join(dir, "ko/tech/post.md"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CollectPosts = %v, want %v", got, want)
	}
	if _, err := CollectPosts("["); err == nil {
		t.Error("CollectPosts with malformed pattern expected error, got nil")
	}
}
