//ff:func feature=quest type=parser control=sequence
//ff:what linkDests 검증 — 본문 prose 전체의 링크 목적지 수집, 코드 펜스 안 링크 구문 제외
package translation

import (
	"fmt"
	"testing"
)

func TestLinkDests(t *testing.T) {
	md := "---\ntitle: x\n---\n\n[a](/x/)\n\n```md\n[b](/y/)\n```\n\n[c](https://e.org)\n"
	got := fmt.Sprint(linkDests(docOf(t, "en", md)))
	if got != "[/x/ https://e.org]" {
		t.Errorf("dests = %s", got)
	}
}
