//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what want 토큰 중 got에 존재하는 비율 계산 — 메타 일치 임계 판정의 입력, want가 비면 1
package gate

// tokenOverlap returns the fraction of want's tokens that appear in got.
func tokenOverlap(want, got []string) float64 {
	if len(want) == 0 {
		return 1
	}
	set := map[string]bool{}
	for _, tok := range got {
		set[tok] = true
	}
	hit := 0
	for _, tok := range want {
		if set[tok] {
			hit++
		}
	}
	return float64(hit) / float64(len(want))
}
