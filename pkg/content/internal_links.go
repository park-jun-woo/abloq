//ff:func feature=content type=parser control=iteration dimension=1
//ff:what 본문 마크다운 링크 중 내부 링크(절대경로 "/" 또는 baseURL 시작) 수 집계 — 클러스터 지표의 입력
package content

import (
	"regexp"
	"strings"
)

var reMDLinkTarget = regexp.MustCompile(`\]\(([^)\s]+)`)

// internalLinks counts markdown link targets pointing into the same site:
// site-absolute paths ("/...") and absolute URLs under baseURL. Relative
// targets (e.g. local images) and external URLs are not counted.
func internalLinks(body, baseURL string) int64 {
	base := strings.TrimRight(baseURL, "/")
	var n int64
	for _, m := range reMDLinkTarget.FindAllStringSubmatch(body, -1) {
		target := m[1]
		if strings.HasPrefix(target, "/") || (base != "" && strings.HasPrefix(target, base)) {
			n++
		}
	}
	return n
}
