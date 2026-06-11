//ff:func feature=insight type=parser control=sequence
//ff:what claims[i]의 진단 라인 번호 — 라인 인덱스 범위 밖이면 1
package insight

// claimLine returns the diagnostic line for claim index i (1 when unknown).
func claimLine(lines []int, i int) int {
	if i < 0 || i >= len(lines) {
		return 1
	}
	return lines[i]
}
