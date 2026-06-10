//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what front matter 원문에서 키 1개의 스칼라 값을 라인 스캔으로 추출 — 따옴표 제거, 없으면 빈 문자열
package gate

import "strings"

// fmLineValue extracts the raw scalar value of a top-level front matter key.
// Line-based on purpose: honest-lastmod compares the literal scalar text.
func fmLineValue(fm, key string) string {
	for _, ln := range strings.Split(fm, "\n") {
		rest, ok := strings.CutPrefix(strings.TrimSpace(ln), key+":")
		if !ok {
			continue
		}
		return strings.Trim(strings.TrimSpace(rest), `"'`)
	}
	return ""
}
