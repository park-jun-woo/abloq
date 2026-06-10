//ff:func feature=visibility type=client control=iteration dimension=1 topic=citation
//ff:what 자기 도메인 인용 추출 — 구조화 인용 URL 목록에서 도메인 부분문자열 매칭으로 자기 글 URL만 선별 (베스트 에포트 v1)
package citation

import "strings"

// Matched filters the engine's citation URLs down to the ones that point at
// our own domain — a best-effort substring heuristic (ExtractorVersion v1):
// trend consistency matters more than precision here.
func Matched(domain string, urls []string) []string {
	if domain == "" {
		return nil
	}
	var out []string
	for _, u := range urls {
		if strings.Contains(u, domain) {
			out = append(out, u)
		}
	}
	return out
}
