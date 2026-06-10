//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what cite 목록 → URL 목록 — 점검기 입력 형태로 사상 (중복 제거는 점검기의 호스트 그룹핑이 담당)
package evidence

// citeURLs projects the probe targets' URLs for the checker.
func citeURLs(cites []cite) []string {
	urls := make([]string, 0, len(cites))
	for _, c := range cites {
		urls = append(urls, c.URL)
	}
	return urls
}
