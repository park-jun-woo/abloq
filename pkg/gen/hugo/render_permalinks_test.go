//ff:func feature=gen type=generator control=sequence
//ff:what renderPermalinks가 섹션 선언 순서 그대로 [permalinks] 블록을 내는지 검증
package hugo

import "testing"

func TestRenderPermalinks(t *testing.T) {
	got := renderPermalinks([]string{"tech", "opinion"})
	want := "\n[permalinks]\ntech = \"/tech/:slug/\"\nopinion = \"/opinion/:slug/\"\n"
	if got != want {
		t.Errorf("renderPermalinks = %q, want %q", got, want)
	}
}
