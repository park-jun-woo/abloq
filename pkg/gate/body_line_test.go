//ff:func feature=gate type=rule control=sequence
//ff:what bodyLine이 본문 인덱스를 파일 라인으로 변환하고 음수에 1을 반환하는지 검증
package gate

import "testing"

func TestBodyLine(t *testing.T) {
	d := &Doc{BodyStart: 7}
	if got := bodyLine(d, 0); got != 7 {
		t.Errorf("bodyLine(0) = %d, want 7", got)
	}
	if got := bodyLine(d, 3); got != 10 {
		t.Errorf("bodyLine(3) = %d, want 10", got)
	}
	if got := bodyLine(d, -1); got != 1 {
		t.Errorf("bodyLine(-1) = %d, want 1", got)
	}
}
