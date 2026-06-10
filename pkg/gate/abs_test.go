//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what abs가 음수/양수/0의 절대값을 반환하는지 검증
package gate

import "testing"

func TestAbs(t *testing.T) {
	if abs(-3) != 3 {
		t.Errorf("abs(-3) = %d, want 3", abs(-3))
	}
	if abs(5) != 5 {
		t.Errorf("abs(5) = %d, want 5", abs(5))
	}
	if abs(0) != 0 {
		t.Errorf("abs(0) = %d, want 0", abs(0))
	}
}
