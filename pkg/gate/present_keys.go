//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 인식 섹션 헤딩 목록을 헤딩 키 집합으로 변환
package gate

// presentKeys returns the set of heading keys present in hits.
func presentKeys(hits []SectionHit) map[string]bool {
	m := map[string]bool{}
	for _, s := range hits {
		m[s.Key] = true
	}
	return m
}
