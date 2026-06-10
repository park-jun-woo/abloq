//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what 두 글의 태그 교집합 크기 — 후보 정렬 1순위 신호
package cluster

import "slices"

// sharedTags counts the tags two articles have in common. Front matter tags
// are short lists, so the quadratic scan is fine and avoids map allocation.
func sharedTags(a, b []string) int64 {
	var n int64
	for _, tag := range a {
		if slices.Contains(b, tag) {
			n++
		}
	}
	return n
}
