//ff:func feature=quest type=rule control=sequence
//ff:what diagsFact 검증 — 첫 진단의 파일:라인/메시지 매핑, 다중 진단은 "(외 N건)" 병기
package writing

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestDiagsFact(t *testing.T) {
	one := []blogyaml.Diagnostic{{File: "content/en/posts/a.md", Line: 7, Rule: "min-sources", Message: "too few"}}
	f := diagsFact("at least 1 source", one)
	if f.Where != "content/en/posts/a.md:7" {
		t.Errorf("Where = %q", f.Where)
	}
	if f.Expected != "at least 1 source" || f.Actual != "too few" {
		t.Errorf("Expected/Actual = %q / %q", f.Expected, f.Actual)
	}
	three := append(one, one[0], one[0])
	f = diagsFact("x", three)
	if !strings.Contains(f.Actual, "(외 2건)") {
		t.Errorf("multi Actual = %q, want 외 2건 suffix", f.Actual)
	}
}
