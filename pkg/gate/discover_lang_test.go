//ff:func feature=gate type=frame control=sequence
//ff:what discoverLang이 선언 섹션의 글만 수집하고 미선언 디렉토리를 무시하는지 검증
package gate

import (
	"path/filepath"
	"testing"
)

func TestDiscoverLang(t *testing.T) {
	b := loadGateBlog(t)
	dir := filepath.Join("testdata", "repo-pass")
	arts := discoverLang(dir, b, buildHeadingIndex(b), "en")
	if len(arts) != 1 {
		t.Fatalf("want 1 article, got %d", len(arts))
	}
	if arts[0].Path != filepath.Join("content", "en", "tech", "hello.md") {
		t.Errorf("path = %s", arts[0].Path)
	}
	if got := discoverLang(dir, b, buildHeadingIndex(b), "fr"); len(got) != 0 {
		t.Errorf("undeclared language dir: want 0 articles, got %d", len(got))
	}
}
