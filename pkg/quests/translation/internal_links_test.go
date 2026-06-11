//ff:func feature=quest type=parser control=sequence
//ff:what internalLinks 검증 — 루트 절대 무확장자 글 링크만 선별, 외부 URL·정적 자산(확장자)·프래그먼트 처리
package translation

import (
	"fmt"
	"testing"
)

func TestInternalLinks(t *testing.T) {
	md := "---\ntitle: x\n---\n\n[a](/posts/p/) [b](https://e.org) [c](/files/doc.pdf) [d](/ko/posts/q/#frag)\n"
	got := fmt.Sprint(internalLinks(docOf(t, "en", md)))
	if got != "[/posts/p/ /ko/posts/q/#frag]" {
		t.Errorf("links = %s", got)
	}
}
