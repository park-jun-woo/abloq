//ff:func feature=gen type=generator control=sequence
//ff:what pinnedGroups가 섹션 미합류 group만 선언 순서대로 자체 헤딩으로 내고, 연속 동일 group은 한 헤딩 아래 모으는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPinnedGroups(t *testing.T) {
	b := &blogyaml.Blog{Site: blogyaml.Site{BaseURL: "https://x.com"}}
	b.Geo.LlmsTxt.Pinned = []blogyaml.LlmsPinned{
		{Title: "Top", URL: "/t"},
		{Title: "A1", URL: "/a1", Group: "Core"},
		{Title: "A2", URL: "/a2", Group: "Core"},
		{Title: "S", URL: "/s", Group: "tech"},
		{Title: "B", URL: "/b", Group: "More"},
	}
	want := "\n## Core\n\n- [A1](https://x.com/a1)\n- [A2](https://x.com/a2)\n" +
		"\n## More\n\n- [B](https://x.com/b)\n"
	if got := pinnedGroups(b, map[string]bool{"tech": true}); got != want {
		t.Errorf("pinnedGroups = %q, want %q", got, want)
	}
	if got := pinnedGroups(&blogyaml.Blog{}, nil); got != "" {
		t.Errorf("no pinned = %q, want empty", got)
	}
}
