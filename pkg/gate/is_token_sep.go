//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what 토큰 구분자 판정 — 공백/구두점/기호 룬은 토큰 경계로 취급
package gate

import "unicode"

// isTokenSep reports whether r separates body tokens.
func isTokenSep(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r)
}
