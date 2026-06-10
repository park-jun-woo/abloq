//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what s[i]에서 시작하는 유효한 %XX 트리플릿의 바이트 값 — '%'가 아니거나 범위 밖·비16진수면 false
package cflog

// tripletVal decodes the %XX triplet starting at s[i]; the second return is
// false when there is no valid triplet there.
func tripletVal(s string, i int) (byte, bool) {
	if s[i] != '%' || i+2 >= len(s) {
		return 0, false
	}
	hi, okHi := hexVal(s[i+1])
	if !okHi {
		return 0, false
	}
	lo, okLo := hexVal(s[i+2])
	if !okLo {
		return 0, false
	}
	return hi<<4 | lo, true
}
