//ff:func feature=gate type=parser control=sequence
//ff:what 텍스트 비교 키 정규화 — NFC + 공백 접기 + 소문자화, 헤딩 분류와 본문 multiset 키가 공유
package gate

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

// normText produces the canonical comparison key for heading and body-line
// matching: NFC + whitespace fold + lowercase.
func normText(s string) string {
	return strings.ToLower(strings.Join(strings.Fields(norm.NFC.String(s)), " "))
}
