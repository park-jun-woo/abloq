//ff:func feature=gen type=generator control=sequence
//ff:what 큐레이션 렌더 검증 — base 단일 스코프 접두 제거, header 블록, pinned(무헤딩/자체 그룹/섹션 합류), 라벨, max_summary 절단, 멱등
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRenderCurated(t *testing.T) {
	b := &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://x.com", Title: "X", Author: "A", DefaultLangInSubdir: true},
		Languages: []string{"en", "ko"},
		Sections:  []string{"opinion", "tech"},
		Geo: blogyaml.Geo{LlmsTxt: blogyaml.LlmsTxtSpec{
			Mode:      "auto",
			Languages: []string{"base"},
			Header:    "Positioning paragraph.\n",
			Pinned: []blogyaml.LlmsPinned{
				{Title: "Master", URL: "/reins.md", Desc: "Index"},
				{Title: "Lead", URL: "/lead/", Group: "Core Content"},
				{Title: "PinOp", URL: "https://x.com/pin/", Group: "Concept"},
			},
			SectionLabels: map[string]string{"opinion": "Concept"},
			MaxSummary:    4,
		}},
	}
	posts := []Post{
		{Lang: "en", Section: "tech", Slug: "t", Title: "T", Date: "2026-01-02"},
		{Lang: "en", Section: "opinion", Slug: "o", Title: "O", Date: "2026-01-01", Description: "abcdefgh"},
	}
	want := "# X\n\n> A — https://x.com\n" +
		"\nPositioning paragraph.\n" +
		"\n- [Master](https://x.com/reins.md): Index\n" +
		"\n## Core Content\n\n- [Lead](https://x.com/lead/)\n" +
		"\n## Concept\n\n- [PinOp](https://x.com/pin/)\n- [O](https://x.com/en/opinion/o/): abcd…\n" +
		"\n## tech\n\n- [T](https://x.com/en/tech/t/)\n"
	if got := string(Render(b, posts)); got != want {
		t.Errorf("Render = %q, want %q", got, want)
	}
	reversed := []Post{posts[1], posts[0]}
	if got := string(Render(b, reversed)); got != want {
		t.Errorf("Render depends on input order: %q", got)
	}
}
