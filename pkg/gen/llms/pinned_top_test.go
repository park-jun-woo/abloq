//ff:func feature=gen type=generator control=sequence
//ff:what pinnedTop이 group 미지정 엔트리만 선언 순서대로 무헤딩 블록으로 내고, 없으면 빈 문자열인지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPinnedTop(t *testing.T) {
	b := &blogyaml.Blog{Site: blogyaml.Site{BaseURL: "https://x.com"}}
	b.Geo.LlmsTxt.Pinned = []blogyaml.LlmsPinned{
		{Title: "A", URL: "/a"},
		{Title: "G", URL: "/g", Group: "Core"},
		{Title: "C", URL: "/c"},
	}
	want := "\n- [A](https://x.com/a)\n- [C](https://x.com/c)\n"
	if got := pinnedTop(b); got != want {
		t.Errorf("pinnedTop = %q, want %q", got, want)
	}
	b.Geo.LlmsTxt.Pinned = []blogyaml.LlmsPinned{{Title: "G", URL: "/g", Group: "Core"}}
	if got := pinnedTop(b); got != "" {
		t.Errorf("grouped-only pinned = %q, want empty", got)
	}
}
