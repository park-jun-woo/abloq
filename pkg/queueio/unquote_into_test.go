//ff:func feature=queueio type=parser control=sequence
//ff:what unquoteInto가 프리픽스 제거+unquote로 필드를 채우고 불량 인용은 에러인지 검증
package queueio

import "testing"

func TestUnquoteInto(t *testing.T) {
	var dst string
	if err := unquoteInto(&dst, "kind: \"refresh\"", "kind: "); err != nil || dst != "refresh" {
		t.Errorf("dst = %q (%v)", dst, err)
	}
	if err := unquoteInto(&dst, "kind: bare", "kind: "); err == nil {
		t.Error("unquoted value must error")
	}
}
