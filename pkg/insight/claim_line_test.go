//ff:func feature=insight type=parser control=sequence
//ff:what claimLine 검증 — 범위 내 인덱스는 해당 라인, 범위 밖·음수는 1
package insight

import "testing"

func TestClaimLine(t *testing.T) {
	lines := []int{4, 9}
	if got := claimLine(lines, 1); got != 9 {
		t.Errorf("want line 9 for index 1, got %d", got)
	}
	if got := claimLine(lines, 2); got != 1 {
		t.Errorf("want fallback 1 for out-of-range index, got %d", got)
	}
	if got := claimLine(lines, -1); got != 1 {
		t.Errorf("want fallback 1 for negative index, got %d", got)
	}
}
