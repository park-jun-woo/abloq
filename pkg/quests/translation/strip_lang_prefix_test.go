//ff:func feature=quest type=parser control=sequence
//ff:what stripLangPrefix 검증 — /{lang}/ 프리픽스 제거와 무프리픽스 링크 통과, 타 언어 프리픽스 비제거
package translation

import "testing"

func TestStripLangPrefix(t *testing.T) {
	if got := stripLangPrefix("/ko/posts/p/", "ko"); got != "/posts/p/" {
		t.Errorf("prefixed = %q", got)
	}
	if got := stripLangPrefix("/posts/p/", "ko"); got != "/posts/p/" {
		t.Errorf("unprefixed = %q", got)
	}
	if got := stripLangPrefix("/ja/posts/p/", "ko"); got != "/ja/posts/p/" {
		t.Errorf("other lang = %q", got)
	}
}
