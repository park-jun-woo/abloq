//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what front matter를 비교용 라인열로 정리 — 우측 공백 제거, lastmod 라인 제외, 말미 빈 라인 제거
package gate

import "strings"

// fmLinesClean prepares one front matter for the intact comparison.
func fmLinesClean(fm string) []string {
	var out []string
	for _, ln := range strings.Split(fm, "\n") {
		t := strings.TrimRight(ln, " \t")
		if strings.HasPrefix(strings.TrimSpace(t), "lastmod:") {
			continue
		}
		out = append(out, t)
	}
	for len(out) > 0 && strings.TrimSpace(out[len(out)-1]) == "" {
		out = out[:len(out)-1]
	}
	return out
}
