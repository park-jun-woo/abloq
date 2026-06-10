//ff:func feature=gate type=rule control=iteration dimension=1 topic=lossless
//ff:what multiset 포함 비교기 — want의 모든 원소가 have에 동일 중복도 이상으로 존재하는지, 첫 누락 원소 반환
package gate

// MultisetSubset reports whether every element of want appears in have with at
// least equal multiplicity. Returns the first missing element when false.
// This is the body-lossless comparator quests reuse as a library.
func MultisetSubset(want, have []string) (string, bool) {
	count := map[string]int{}
	for _, h := range have {
		count[h]++
	}
	for _, w := range want {
		if count[w] <= 0 {
			return w, false
		}
		count[w]--
	}
	return "", true
}
