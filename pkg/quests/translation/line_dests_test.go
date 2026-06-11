//ff:func feature=quest type=parser control=sequence
//ff:what lineDests 검증 — 라인 1개의 비이미지 링크 목적지 추출, 이미지 매치와 제목 따옴표 뒤 잔여 무시
package translation

import (
	"fmt"
	"testing"
)

func TestLineDests(t *testing.T) {
	got := fmt.Sprint(lineDests(`see [a](/x/) and ![img](/i.png) and [b](https://e.org "t")`))
	if got != "[/x/ https://e.org]" {
		t.Errorf("dests = %s", got)
	}
	if d := lineDests("no links here"); d != nil {
		t.Errorf("plain line: %v", d)
	}
}
