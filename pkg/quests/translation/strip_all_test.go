//ff:func feature=quest type=parser control=sequence
//ff:what stripAll 검증 — 링크 목록 전체에 프리픽스 제거 적용
package translation

import (
	"fmt"
	"testing"
)

func TestStripAll(t *testing.T) {
	got := fmt.Sprint(stripAll([]string{"/ko/a/", "/b/"}, "ko"))
	if got != "[/a/ /b/]" {
		t.Errorf("stripped = %s", got)
	}
}
