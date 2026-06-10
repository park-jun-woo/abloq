//ff:func feature=gen type=generator control=sequence
//ff:what llms.txt 렌더가 입력 순서와 무관하게 같은 바이트를 내고 그룹 헤더·글 줄을 정확히 배치하는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRender(t *testing.T) {
	b := &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://x.com", Title: "X Blog", Author: "A"},
		Languages: []string{"ko", "en"},
		Sections:  []string{"opinion"},
	}
	posts := []Post{
		{Lang: "en", Section: "opinion", Slug: "e", Title: "E", Date: "2026-01-01"},
		{Lang: "ko", Section: "opinion", Slug: "k-old", Title: "KO", Date: "2026-01-01"},
		{Lang: "ko", Section: "opinion", Slug: "k-new", Title: "KN", Date: "2026-02-01", Description: "d"},
	}
	want := "# X Blog\n\n> A — https://x.com\n" +
		"\n## ko/opinion\n\n" +
		"- [KN](https://x.com/ko/opinion/k-new/): d\n" +
		"- [KO](https://x.com/ko/opinion/k-old/)\n" +
		"\n## en/opinion\n\n" +
		"- [E](https://x.com/en/opinion/e/)\n"
	if got := string(Render(b, posts)); got != want {
		t.Errorf("Render = %q, want %q", got, want)
	}
	reversed := []Post{posts[2], posts[1], posts[0]}
	if got := string(Render(b, reversed)); got != want {
		t.Errorf("Render depends on input order: %q", got)
	}
}
