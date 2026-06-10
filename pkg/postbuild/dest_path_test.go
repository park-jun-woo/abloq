//ff:func feature=postbuild type=generator control=iteration dimension=1
//ff:what DestPath가 단일 파일과 번들 소스를 public/{lang}/{section}/{slug}.md로 변환하는지 검증
package postbuild

import (
	"path/filepath"
	"testing"
)

func TestDestPath(t *testing.T) {
	cases := []struct{ src, want string }{
		{"content/ko/tech/post.md", "public/ko/tech/post.md"},
		{"content/ko/tech/bundle/index.md", "public/ko/tech/bundle.md"},
	}
	for _, c := range cases {
		got := DestPath("content", "public", c.src)
		if got != filepath.FromSlash(c.want) {
			t.Errorf("DestPath(%s) = %s, want %s", c.src, got, c.want)
		}
	}
	if got := DestPath("content", "public", "/abs/post.md"); got != filepath.FromSlash("public/abs/post.md") {
		t.Errorf("DestPath rel-failure fallback = %s, want public/abs/post.md", got)
	}
}
