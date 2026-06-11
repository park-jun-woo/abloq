//ff:func feature=gen type=generator control=sequence
//ff:what sectionPinned가 그룹 헤딩 텍스트와 group이 일치하는 pinned만 선언 순서대로 내고, 없으면 빈 문자열인지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestSectionPinned(t *testing.T) {
	b := &blogyaml.Blog{Site: blogyaml.Site{BaseURL: "https://x.com"}}
	b.Geo.LlmsTxt.Pinned = []blogyaml.LlmsPinned{
		{Title: "A", URL: "/a", Group: "Concept"},
		{Title: "B", URL: "/b", Group: "tech"},
		{Title: "C", URL: "/c", Group: "Concept"},
	}
	want := "- [A](https://x.com/a)\n- [C](https://x.com/c)\n"
	if got := sectionPinned(b, "Concept"); got != want {
		t.Errorf("sectionPinned(Concept) = %q, want %q", got, want)
	}
	if got := sectionPinned(b, "none"); got != "" {
		t.Errorf("sectionPinned(none) = %q, want empty", got)
	}
}
