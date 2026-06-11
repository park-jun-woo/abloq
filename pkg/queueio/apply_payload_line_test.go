//ff:func feature=queueio type=parser control=sequence
//ff:what applyPayloadLine이 payload 항목을 unquote 적재하고 콜론 부재·불량 인용은 에러인지 검증
package queueio

import "testing"

func TestApplyPayloadLine(t *testing.T) {
	it := Item{Payload: map[string]string{}}
	if err := applyPayloadLine(&it, "  claims: \"[1,2]\""); err != nil {
		t.Fatalf("applyPayloadLine: %v", err)
	}
	if it.Payload["claims"] != "[1,2]" {
		t.Errorf("payload = %v", it.Payload)
	}
	if err := applyPayloadLine(&it, "  no-colon-here"); err == nil {
		t.Error("missing ': ' separator must error")
	}
	if err := applyPayloadLine(&it, "  k: not-quoted"); err == nil {
		t.Error("unquoted value must error")
	}
}
