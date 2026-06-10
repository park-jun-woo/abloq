//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what splitFrontMatter가 정상 블록/블록 없음/미종결 블록/CRLF를 처리하는지 검증
package gate

import "testing"

func TestSplitFrontMatter(t *testing.T) {
	cases := []struct {
		name, in, wantFM, wantBody string
		wantOK                     bool
	}{
		{"well-formed", "---\ntitle: x\n---\nbody\n", "title: x", "body\n", true},
		{"crlf", "---\r\ntitle: x\r\n---\r\nbody\r\n", "title: x", "body\n", true},
		{"no front matter", "body only\n", "", "body only\n", false},
		{"unterminated", "---\ntitle: x\nbody\n", "", "---\ntitle: x\nbody\n", false},
		{"fence not first", "x\n---\ntitle\n---\n", "", "x\n---\ntitle\n---\n", false},
		{"no newline after fence", "---\ntitle: x\n---", "title: x", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkSplitFrontMatter(t, tc.in, tc.wantFM, tc.wantBody, tc.wantOK) })
	}
}
