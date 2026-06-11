//ff:func feature=quest type=parser control=sequence
//ff:what sectionKeys 검증 — 인식된 섹션 헤딩 키의 문서 순서 추출, 자유 헤딩은 미포함
package translation

import (
	"fmt"
	"testing"
)

func TestSectionKeys(t *testing.T) {
	origin, _ := passPair()
	got := fmt.Sprint(sectionKeys(docOf(t, "en", origin)))
	if got != "[sources]" {
		t.Errorf("keys = %s, want [sources] (## Setup is free-form)", got)
	}
}
