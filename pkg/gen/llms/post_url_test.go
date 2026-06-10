//ff:func feature=gen type=generator control=sequence
//ff:what postURL이 baseURL 끝 슬래시를 정리하고 /언어/섹션/slug/ 정규 URL을 만드는지, 루트 서빙 기본 언어는 언어 세그먼트를 생략하는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPostURL(t *testing.T) {
	b := &blogyaml.Blog{Languages: []string{"en", "ko"}}
	b.Site.BaseURL = "https://x.example.com"
	b.Site.DefaultLangInSubdir = true
	p := Post{Lang: "ko", Section: "tech", Slug: "hello"}
	want := "https://x.example.com/ko/tech/hello/"
	if got := postURL(b, p); got != want {
		t.Errorf("postURL = %q, want %q", got, want)
	}
	b.Site.BaseURL = "https://x.example.com/"
	if got := postURL(b, p); got != want {
		t.Errorf("postURL with trailing slash = %q, want %q", got, want)
	}
	b.Site.DefaultLangInSubdir = false
	if got := postURL(b, p); got != want {
		t.Errorf("postURL non-default lang at root site = %q, want %q", got, want)
	}
	en := Post{Lang: "en", Section: "tech", Slug: "hello"}
	wantRoot := "https://x.example.com/tech/hello/"
	if got := postURL(b, en); got != wantRoot {
		t.Errorf("postURL root-served default lang = %q, want %q", got, wantRoot)
	}
}
