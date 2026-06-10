//ff:func feature=gen type=generator control=sequence
//ff:what postLine이 "- [제목](URL)" 형식을 만들고 설명이 있을 때만 ": 설명"을 붙이는지 검증
package llms

import "testing"

func TestPostLine(t *testing.T) {
	p := Post{Lang: "ko", Section: "tech", Slug: "a", Title: "A"}
	want := "- [A](https://x.com/ko/tech/a/)"
	if got := postLine("https://x.com", p); got != want {
		t.Errorf("postLine = %q, want %q", got, want)
	}
	p.Description = "desc"
	if got := postLine("https://x.com", p); got != want+": desc" {
		t.Errorf("postLine with description = %q, want %q", got, want+": desc")
	}
}
