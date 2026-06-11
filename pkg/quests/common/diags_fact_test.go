//ff:func feature=quest type=rule control=sequence
//ff:what DiagsFact 검증 — 첫 진단의 파일:라인/메시지 매핑과 다중 진단의 "(외 N건)" 병기
package common

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestDiagsFact(t *testing.T) {
	diags := []blogyaml.Diagnostic{
		{File: "content/en/posts/a.md", Line: 3, Rule: "x", Message: "first"},
		{File: "content/en/posts/a.md", Line: 9, Rule: "x", Message: "second"},
	}
	f := DiagsFact("expected text", diags)
	if f.Where != "content/en/posts/a.md:3" {
		t.Errorf("Where = %q", f.Where)
	}
	if f.Expected != "expected text" || f.Actual != "first (외 1건)" {
		t.Errorf("Expected/Actual = %q / %q", f.Expected, f.Actual)
	}
	if got := DiagsFact("e", diags[:1]).Actual; got != "first" {
		t.Errorf("single diag Actual = %q", got)
	}
}
