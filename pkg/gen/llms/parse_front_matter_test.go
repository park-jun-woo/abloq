//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what parseFrontMatter가 정상 블록을 디코드하고 부재·미종결·YAML 오류를 false로 거르는지 검증
package llms

import "testing"

func TestParseFrontMatter(t *testing.T) {
	cases := []struct {
		name      string
		data      string
		wantOK    bool
		wantTitle string
		wantDraft bool
	}{
		{"valid", "---\ntitle: Hello\ndate: 2026-01-02\ndraft: false\n---\nbody\n", true, "Hello", false},
		{"draft true", "---\ntitle: D\ndraft: true\n---\n", true, "D", true},
		{"unknown keys ignored", "---\ntitle: T\ntags: [a, b]\n---\n", true, "T", false},
		{"delimiter at eof", "---\ntitle: E\n---", true, "E", false},
		{"no front matter", "body only\n", false, "", false},
		{"unterminated", "---\ntitle: X\n", false, "", false},
		{"invalid yaml", "---\ntitle: [\n---\n", false, "", false},
		{"empty file", "", false, "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkParseFrontMatter(t, tc.data, tc.wantOK, tc.wantTitle, tc.wantDraft) })
	}
}
