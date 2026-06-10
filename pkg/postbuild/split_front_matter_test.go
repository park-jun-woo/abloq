//ff:func feature=postbuild type=parser control=iteration dimension=1
//ff:what splitFrontMatter가 정상 블록 분리, 블록 없음, CRLF 정규화, 미종결 블록을 처리하는지 검증
package postbuild

import "testing"

func TestSplitFrontMatter(t *testing.T) {
	cases := []struct {
		name, in, fm, body string
	}{
		{"normal", "---\ntitle: A\n---\nbody\n", "title: A", "body\n"},
		{"none", "just body\n", "", "just body\n"},
		{"crlf", "---\r\ntitle: A\r\n---\r\nbody\r\n", "title: A", "body\n"},
		{"unterminated", "---\ntitle: A\n", "", "---\ntitle: A\n"},
		{"no-trailing-newline", "---\ntitle: A\n---", "title: A", ""},
	}
	for _, c := range cases {
		fm, body := splitFrontMatter(c.in)
		if fm != c.fm || body != c.body {
			t.Errorf("%s: splitFrontMatter = (%q, %q), want (%q, %q)", c.name, fm, body, c.fm, c.body)
		}
	}
}
