//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what Render가 6개 섹션 헤더를 고정 순서로 포함하고 멱등(바이트 동일)인지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	b := testBlog()
	out := string(Render(b))
	if again := string(Render(b)); again != out {
		t.Fatal("Render is not idempotent")
	}
	wants := []string{
		"# CLAUDE.md — T Blog 운영 매뉴얼",
		"## 블로그",
		"## 디렉토리 규약",
		"## 글 구조 계약",
		"## 게시 절차",
		"## 명령어",
		"## 금지 사항",
	}
	last := -1
	for _, w := range wants {
		idx := strings.Index(out, w)
		if idx < 0 {
			t.Errorf("want %q in CLAUDE.md", w)
			continue
		}
		if idx < last {
			t.Errorf("%q out of order", w)
		}
		last = idx
	}
}
