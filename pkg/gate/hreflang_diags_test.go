//ff:func feature=gate type=rule control=sequence
//ff:what hreflangDiags가 누락 언어별 진단·미빌드 페이지 진단을 만들고 루트 서빙 기본 언어의 페이지 경로를 푸는지 검증
package gate

import (
	"path/filepath"
	"testing"
)

func TestHreflangDiags(t *testing.T) {
	b := loadGateBlog(t)
	en := artFromMD(t, b, "en", "tech", "hello", "repo-pass/content/en/tech/hello.md")
	passDir := filepath.Join("testdata", "repo-pass")
	if diags := hreflangDiags(passDir, b, en, []string{"ko", "en"}); len(diags) != 0 {
		t.Fatalf("complete page: want 0, got %v", diags)
	}
	failDir := filepath.Join("testdata", "repo-hreflang-fail")
	diags := hreflangDiags(failDir, b, en, []string{"ko", "en"})
	checkDiags(t, diags, 1, "hreflang-complete", "lacks hreflang alternate for ko")
	ko := artFromMD(t, b, "ko", "tech", "hello", "repo-pass/content/ko/tech/hello.md")
	missing := hreflangDiags(failDir, b, ko, []string{"ko", "en"})
	checkDiags(t, missing, 1, "hreflang-complete", "missing — run the site build")

	root := *b
	root.Site.DefaultLangInSubdir = false
	rootKo := artFromMD(t, b, "ko", "tech", "hello", "repo-pass/content/ko/tech/hello.md")
	got := hreflangDiags(passDir, &root, rootKo, []string{"ko"})
	checkDiags(t, got, 1, "hreflang-complete", filepath.Join("public", "tech", "hello", "index.html")+" missing")
}
