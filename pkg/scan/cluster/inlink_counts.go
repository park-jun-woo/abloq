//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what 글별 인링크 수 집계 — 확정 아웃링크의 역방향 합, no-isolated-post 판정(0 = 고립)의 입력
package cluster

// inlinkCounts derives each node's in-degree from the finalized outlinks.
// Outlinks are deduplicated per article, so one article contributes at most
// one inlink to another.
func inlinkCounts(posts []post) map[string]int64 {
	counts := map[string]int64{}
	for _, p := range posts {
		for _, key := range p.Outlinks {
			counts[key]++
		}
	}
	return counts
}
