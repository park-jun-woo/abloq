//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what publishSection이 기본 언어 경로/게이트/번역/generate/postbuild 순서의 게시 절차를 포함하는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestPublishSection(t *testing.T) {
	out := publishSection(testBlog())
	wants := []string{
		"content/ko/{section}/{slug}.md",
		"abloq gate --offline .",
		"동일 slug",
		"abloq generate .",
		"abloq postbuild md .",
		"abloq image convert",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in publish section, got:\n%s", w, out)
		}
	}
}
