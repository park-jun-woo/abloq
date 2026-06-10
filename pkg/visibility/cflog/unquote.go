//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 퍼센트 인코딩 디코드 — %XX만 풀고 비정상 시퀀스는 원문 유지, '+'는 건드리지 않음 (python urllib unquote 동등)
//ff:why analyze-stats.py L59-61이 urllib.parse.unquote로 URI·UA를 디코드한다 — net/url의 PathUnescape는 비정상 %에서 에러라 의미가 다르고, 대조 일치를 위해 관용 디코드가 필요하다 (Phase012)
package cflog

// unquote percent-decodes s the way python's urllib.parse.unquote does:
// every valid %XX triplet becomes its byte, invalid sequences stay verbatim,
// and '+' is left alone (that is unquote_plus, not unquote).
func unquote(s string) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if b, ok := tripletVal(s, i); ok {
			out = append(out, b)
			i += 2
			continue
		}
		out = append(out, s[i])
	}
	return string(out)
}
