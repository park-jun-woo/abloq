//ff:func feature=gen type=generator control=sequence
//ff:what postGroups가 그룹 전환마다 헤딩과 합류 pinned 선두 줄을 내고, 같은 그룹 글은 헤딩 없이 이어 붙이는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPostGroups(t *testing.T) {
	b := &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://x.com", DefaultLangInSubdir: true},
		Languages: []string{"ko"},
		Sections:  []string{"opinion", "tech"},
	}
	b.Geo.LlmsTxt.SectionLabels = map[string]string{"opinion": "Concept"}
	b.Geo.LlmsTxt.Pinned = []blogyaml.LlmsPinned{{Title: "Pin", URL: "/p/", Group: "Concept"}}
	sorted := []Post{
		{Lang: "ko", Section: "opinion", Slug: "o", Title: "O"},
		{Lang: "ko", Section: "tech", Slug: "t", Title: "T"},
		{Lang: "ko", Section: "tech", Slug: "t2", Title: "T2"},
	}
	want := "\n## Concept\n\n- [Pin](https://x.com/p/)\n- [O](https://x.com/ko/opinion/o/)\n" +
		"\n## tech\n\n- [T](https://x.com/ko/tech/t/)\n- [T2](https://x.com/ko/tech/t2/)\n"
	if got := postGroups(b, sorted, false); got != want {
		t.Errorf("postGroups = %q, want %q", got, want)
	}
	multi := postGroups(b, sorted[:1], true)
	wantMulti := "\n## ko/Concept\n\n- [O](https://x.com/ko/opinion/o/)\n"
	if multi != wantMulti {
		t.Errorf("postGroups multi = %q, want %q", multi, wantMulti)
	}
}
