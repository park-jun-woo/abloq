//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what 두 발행일의 거리(일수, 절대값) — 후보 정렬 3순위 신호(근접 우선)
package cluster

// dateDistance measures how many days apart two articles were published.
func dateDistance(a, b string) int64 {
	d := epochDay(a) - epochDay(b)
	if d < 0 {
		return -d
	}
	return d
}
