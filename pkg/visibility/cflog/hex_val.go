//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 16진수 문자 1개를 값으로 — 0-9/a-f/A-F 외에는 false
package cflog

// hexVal decodes one hex digit; the second return is false for non-hex.
func hexVal(c byte) (byte, bool) {
	if c >= '0' && c <= '9' {
		return c - '0', true
	}
	if c >= 'a' && c <= 'f' {
		return c - 'a' + 10, true
	}
	if c >= 'A' && c <= 'F' {
		return c - 'A' + 10, true
	}
	return 0, false
}
