//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what 이전 점검 상태 → 키별 연속 실패 수 맵 — 연속 카운트 갱신의 조회 인덱스
package evidence

// prevFailures indexes the previous scan's consecutive-failure counts by the
// citation_checks key.
func prevFailures(prev []Check) map[string]int64 {
	m := make(map[string]int64, len(prev))
	for _, c := range prev {
		m[checkKey(c.URL, c.Lang, c.Section, c.Slug)] = c.ConsecutiveFailures
	}
	return m
}
