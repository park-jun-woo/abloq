//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what 본문을 공백·구두점 정규화 토큰열로 변환 — honest-lastmod의 사소-diff 게임 방지 입력
package gate

import "strings"

// Tokens normalizes a body into comparison tokens: lowercase, with whitespace,
// punctuation and symbols all treated as separators.
func Tokens(body string) []string {
	return strings.FieldsFunc(strings.ToLower(body), isTokenSep)
}
