//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what headerSection이 제목/저자/baseURL/기본 언어 표시/섹션 목록을 포함하는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestHeaderSection(t *testing.T) {
	out := headerSection(testBlog())
	wants := []string{
		"T Blog 운영 매뉴얼",
		"baseURL: https://t.example.com",
		"저자: Tester",
		`기본 언어는 "ko"`,
		"섹션: opinion, tech",
		"abloq claudemd",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in header section, got:\n%s", w, out)
		}
	}
}
