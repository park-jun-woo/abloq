//ff:func feature=gate type=rule control=sequence
//ff:what [hreflang-complete] 완전한 상호 참조 PASS, 대체 링크 누락·페이지 미빌드 FAIL, public 없음 스킵 검증
package gate

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRuleHreflangComplete(t *testing.T) {
	passDir := filepath.Join("testdata", "repo-pass")
	b, _, err := blogyaml.Load(filepath.Join(passDir, "blog.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	pass := NewTarget(passDir, b, Discover(passDir, b))
	if diags := ruleHreflangComplete(pass); len(diags) != 0 {
		t.Fatalf("repo-pass: want 0 diagnostics, got %v", diags)
	}

	failDir := filepath.Join("testdata", "repo-hreflang-fail")
	fail := NewTarget(failDir, b, Discover(failDir, b))
	diags := ruleHreflangComplete(fail)
	checkDiags(t, diags, 2, "hreflang-complete", "missing — run the site build")

	noBuild := NewTarget("testdata", b, Discover(passDir, b))
	if diags := ruleHreflangComplete(noBuild); len(diags) != 0 {
		t.Errorf("no public dir: want skip, got %v", diags)
	}
}
