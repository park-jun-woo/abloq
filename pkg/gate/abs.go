//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what 정수 절대값 — 토큰 대칭차 합산용
package gate

// abs returns the absolute value of n.
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
