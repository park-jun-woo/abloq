//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what numClaimDiags 케이스 — 무출처 주장 2건이 각자의 파일 라인으로 진단되는지 검증
package gate

import "testing"

func TestNumClaimDiags(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromContent(t, b, "Throughput improved by 42%.\n\nLatency dropped to 12ms.\n")
	diags := numClaimDiags(a)
	if len(diags) != 2 {
		t.Fatalf("want 2 diagnostics, got %v", diags)
	}
	if diags[0].Line != 1 || diags[1].Line != 3 {
		t.Errorf("want lines 1 and 3, got %d and %d", diags[0].Line, diags[1].Line)
	}
}
