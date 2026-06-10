//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what 두 토큰열의 multiset 대칭차 크기 계산 — honest-lastmod의 본문 실변경량 측정
package gate

// TokenDiff returns the size of the multiset symmetric difference between two
// token sequences: the number of tokens added plus the number removed.
func TokenDiff(a, b []string) int {
	count := map[string]int{}
	for _, t := range a {
		count[t]++
	}
	for _, t := range b {
		count[t]--
	}
	diff := 0
	for _, n := range count {
		diff += abs(n)
	}
	return diff
}
