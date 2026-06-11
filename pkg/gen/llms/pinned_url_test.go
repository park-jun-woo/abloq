//ff:func feature=gen type=generator control=sequence
//ff:what pinnedURL이 "/" 시작 경로를 baseURL(trailing 슬래시 제거)과 결합하고 절대 URL은 그대로 두는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPinnedURL(t *testing.T) {
	b := &blogyaml.Blog{Site: blogyaml.Site{BaseURL: "https://x.com/"}}
	if got := pinnedURL(b, "/reins.md"); got != "https://x.com/reins.md" {
		t.Errorf("site-rooted path = %q, want baseURL joined without double slash", got)
	}
	if got := pinnedURL(b, "https://e.com/p"); got != "https://e.com/p" {
		t.Errorf("absolute url = %q, want unchanged", got)
	}
}
