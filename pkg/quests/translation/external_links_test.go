//ff:func feature=quest type=parser control=sequence
//ff:what externalLinks 검증 — http/https 목적지만 선별, 내부 경로 제외
package translation

import (
	"fmt"
	"testing"
)

func TestExternalLinks(t *testing.T) {
	md := "---\ntitle: x\n---\n\n[a](/x/) [b](https://e.org/p) [c](http://e.org/q)\n"
	got := fmt.Sprint(externalLinks(docOf(t, "en", md)))
	if got != "[https://e.org/p http://e.org/q]" {
		t.Errorf("urls = %s", got)
	}
}
