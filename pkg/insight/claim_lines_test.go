//ff:func feature=insight type=parser control=sequence
//ff:what claims 시퀀스 항목별 라인 수집 검증 — 정상 2건, claims 없음, 비매핑 루트, 깨진 yaml
package insight

import (
	"reflect"
	"testing"
)

func TestClaimLines(t *testing.T) {
	data := []byte("topic: t\nclaims:\n  - id: a\n    text: x\n  - id: b\n    text: y\n")
	if got := claimLines(data); !reflect.DeepEqual(got, []int{3, 5}) {
		t.Errorf("want claim lines [3 5], got %v", got)
	}
	if got := claimLines([]byte("topic: t\n")); got != nil {
		t.Errorf("want nil without claims, got %v", got)
	}
	if got := claimLines([]byte("- a\n- b\n")); got != nil {
		t.Errorf("want nil for non-mapping root, got %v", got)
	}
	if got := claimLines([]byte(":\n  :bad")); got != nil {
		t.Errorf("want nil for broken yaml, got %v", got)
	}
}
