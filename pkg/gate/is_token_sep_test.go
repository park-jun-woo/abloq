//ff:func feature=gate type=parser control=iteration dimension=1 topic=baseline
//ff:what isTokenSep이 공백/구두점/기호를 경계로, 문자·숫자를 토큰으로 판정하는지 검증
package gate

import "testing"

func TestIsTokenSep(t *testing.T) {
	cases := []struct {
		name string
		r    rune
		want bool
	}{
		{"space", ' ', true},
		{"newline", '\n', true},
		{"comma", ',', true},
		{"percent sign", '%', true},
		{"plus symbol", '+', true},
		{"letter", 'a', false},
		{"hangul", '한', false},
		{"digit", '7', false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isTokenSep(tc.r); got != tc.want {
				t.Errorf("isTokenSep(%q) = %v, want %v", tc.r, got, tc.want)
			}
		})
	}
}
