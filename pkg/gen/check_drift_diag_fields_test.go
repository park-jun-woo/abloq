//ff:func feature=gen type=rule control=sequence topic=drift
//ff:what 드리프트 진단 1건의 파일/라인/룰ID/메시지 내용을 기대값과 비교 검증
package gen

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func checkDriftDiagFields(t *testing.T, d blogyaml.Diagnostic, wantFile string, wantLine int, wantRule, wantInMsg string) {
	t.Helper()
	if d.File != wantFile {
		t.Errorf("diag file = %q, want %q", d.File, wantFile)
	}
	if d.Line != wantLine {
		t.Errorf("diag line = %d, want %d", d.Line, wantLine)
	}
	if d.Rule != wantRule {
		t.Errorf("diag rule = %q, want %q", d.Rule, wantRule)
	}
	if !strings.Contains(d.Message, wantInMsg) {
		t.Errorf("diag message = %q, want it to contain %q", d.Message, wantInMsg)
	}
}
