//ff:func feature=content type=parser control=sequence
//ff:what IndexRepo가 blog.yaml 부재는 IO 에러로, 스키마 위반은 "blog.yaml invalid" 진단 에러로 구분해 실패하는지 검증
package content

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestIndexRepoInvalid(t *testing.T) {
	if _, err := IndexRepo(t.TempDir()); err == nil {
		t.Fatal("missing blog.yaml must be an error")
	}
	dir := t.TempDir()
	bad := "site:\n  baseURL: not-a-url\nlanguages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(bad), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := IndexRepo(dir)
	if err == nil || !strings.Contains(err.Error(), "blog.yaml invalid") {
		t.Fatalf("schema violation must surface as a diagnostic error, got %v", err)
	}
}
