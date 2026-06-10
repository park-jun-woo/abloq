//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what front matter 키가 위치한 파일 라인 번호 계산 — 펜스 1줄 보정, 없으면 1
package gate

import "strings"

// fmKeyLine returns the 1-based file line of a top-level front matter key
// (the opening fence is line 1), or 1 when the key is absent.
func fmKeyLine(fm, key string) int {
	for i, ln := range strings.Split(fm, "\n") {
		if strings.HasPrefix(strings.TrimSpace(ln), key+":") {
			return i + 2
		}
	}
	return 1
}
