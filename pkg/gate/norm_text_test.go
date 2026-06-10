//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what normText가 공백 접기/소문자화/NFC 정규화로 동일 비교 키를 만드는지 검증
package gate

import "testing"

func TestNormText(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"fold spaces", "  Further   Reading ", "further reading"},
		{"lowercase", "Related Posts", "related posts"},
		{"korean", "관련 글", "관련 글"},
		{"nfc", "étude", "étude"},
		{"empty", "   ", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := normText(tc.in); got != tc.want {
				t.Errorf("normText(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
