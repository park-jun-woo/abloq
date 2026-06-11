//ff:func feature=gen type=generator control=sequence
//ff:what pinnedLine이 "- [제목](URL)"을 만들고 desc가 있을 때만 ": 설명"을 붙이며, "/" 시작 url은 baseURL과 결합하는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPinnedLine(t *testing.T) {
	b := &blogyaml.Blog{Site: blogyaml.Site{BaseURL: "https://x.com"}}
	p := blogyaml.LlmsPinned{Title: "Master", URL: "/reins.md"}
	if got := pinnedLine(b, p); got != "- [Master](https://x.com/reins.md)" {
		t.Errorf("pinnedLine = %q, want site-rooted url joined without desc", got)
	}
	p.Desc = "Index"
	if got := pinnedLine(b, p); got != "- [Master](https://x.com/reins.md): Index" {
		t.Errorf("pinnedLine with desc = %q, want trailing ': Index'", got)
	}
	abs := blogyaml.LlmsPinned{Title: "Ext", URL: "https://e.com/p"}
	if got := pinnedLine(b, abs); got != "- [Ext](https://e.com/p)" {
		t.Errorf("pinnedLine absolute = %q, want url unchanged", got)
	}
}
