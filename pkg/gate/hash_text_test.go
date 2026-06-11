//ff:func feature=gate type=generator control=sequence topic=evidence
//ff:what HashText가 결정적 16자 hex를 내고 텍스트가 다르면 키도 다른지 검증
package gate

import "testing"

func TestHashText(t *testing.T) {
	a := HashText("처리량이 40% 증가했다.")
	if len(a) != 16 {
		t.Fatalf("hash length = %d, want 16", len(a))
	}
	if a != HashText("처리량이 40% 증가했다.") {
		t.Error("hash must be deterministic")
	}
	if a == HashText("지연이 120ms 단축됐다.") {
		t.Error("different texts must not collide")
	}
}
