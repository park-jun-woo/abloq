//ff:func feature=postbuild type=generator control=iteration dimension=1
//ff:what DestPath가 단일 파일과 번들 소스를 public/{lang}/{section}/{slug}.md로 변환하는지, 루트 서빙 기본 언어의 언어 디렉토리를 생략하는지 검증
package postbuild

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestDestPath(t *testing.T) {
	sub := &blogyaml.Blog{Languages: []string{"ko", "en"}}
	sub.Site.DefaultLangInSubdir = true
	cases := []struct{ src, want string }{
		{"content/ko/tech/post.md", "public/ko/tech/post.md"},
		{"content/ko/tech/bundle/index.md", "public/ko/tech/bundle.md"},
	}
	for _, c := range cases {
		got := DestPath("content", "public", c.src, sub)
		if got != filepath.FromSlash(c.want) {
			t.Errorf("DestPath(%s) = %s, want %s", c.src, got, c.want)
		}
	}
	if got := DestPath("content", "public", "/abs/post.md", sub); got != filepath.FromSlash("public/abs/post.md") {
		t.Errorf("DestPath rel-failure fallback = %s, want public/abs/post.md", got)
	}
	root := &blogyaml.Blog{Languages: []string{"en", "ko"}}
	root.Site.DefaultLangInSubdir = false
	if got := DestPath("content", "public", "content/en/tech/post.md", root); got != filepath.FromSlash("public/tech/post.md") {
		t.Errorf("DestPath root-served default lang = %s, want public/tech/post.md", got)
	}
	if got := DestPath("content", "public", "content/ko/tech/post.md", root); got != filepath.FromSlash("public/ko/tech/post.md") {
		t.Errorf("DestPath other lang on root site = %s, want public/ko/tech/post.md", got)
	}
}
