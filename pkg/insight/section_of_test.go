//ff:func feature=insight type=parser control=iteration dimension=1
//ff:what 글 경로 → 섹션 도출 검증 — 플랫·번들·중첩 content·content 밖 경로
package insight

import (
	"path/filepath"
	"testing"
)

func TestSectionOf(t *testing.T) {
	cases := []struct {
		path string
		want string
	}{
		{filepath.Join("content", "en", "tech", "post.md"), "tech"},
		{filepath.Join("content", "en", "opinion", "post", "index.md"), "opinion"},
		{filepath.Join("testdata", "content", "ko", "tech", "post.md"), "tech"},
		{filepath.Join("repo", "content", "x", "content", "en", "writing", "p.md"), "writing"},
		{filepath.Join("elsewhere", "post.md"), ""},
		{filepath.Join("content", "en"), ""},
	}
	for _, c := range cases {
		if got := sectionOf(c.path); got != c.want {
			t.Errorf("sectionOf(%q): want %q, got %q", c.path, c.want, got)
		}
	}
}
