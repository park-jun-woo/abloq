//ff:func feature=gen type=generator control=sequence
//ff:what postURL이 baseURL 끝 슬래시를 정리하고 /언어/섹션/slug/ 정규 URL을 만드는지 검증
package llms

import "testing"

func TestPostURL(t *testing.T) {
	p := Post{Lang: "ko", Section: "tech", Slug: "hello"}
	want := "https://x.example.com/ko/tech/hello/"
	if got := postURL("https://x.example.com", p); got != want {
		t.Errorf("postURL = %q, want %q", got, want)
	}
	if got := postURL("https://x.example.com/", p); got != want {
		t.Errorf("postURL with trailing slash = %q, want %q", got, want)
	}
}
