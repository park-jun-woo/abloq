//ff:func feature=scan type=client control=iteration dimension=1 topic=evidence
//ff:what 호스트별 점검 고루틴 n개의 부분 결과 맵을 하나의 url→판정 맵으로 합침
package evidence

// mergeResults drains n partial maps from the host workers into one result.
func mergeResults(out chan map[string]string, n int) map[string]string {
	res := map[string]string{}
	for range n {
		for u, st := range <-out {
			res[u] = st
		}
	}
	return res
}
