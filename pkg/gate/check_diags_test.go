//ff:func feature=gate type=rule control=sequence topic=diagnostics
//ff:what 룰 실행 결과 검증 헬퍼 — 진단 수, 첫 진단의 룰ID와 메시지 부분 일치를 확인
package gate

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func checkDiags(t *testing.T, diags []blogyaml.Diagnostic, wantCount int, wantRule, wantMsgPart string) {
	t.Helper()
	if len(diags) != wantCount {
		t.Fatalf("want %d diagnostics, got %d: %v", wantCount, len(diags), diags)
	}
	if wantCount == 0 {
		return
	}
	if diags[0].Rule != wantRule {
		t.Errorf("want rule %s, got %q", wantRule, diags[0].Rule)
	}
	if !strings.Contains(diags[0].Message, wantMsgPart) {
		t.Errorf("want message containing %q, got %q", wantMsgPart, diags[0].Message)
	}
}
