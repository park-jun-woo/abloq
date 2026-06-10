//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what 그래프 노드 키 집합 조립 — 아웃링크 코퍼스 실재 필터의 조회 대상
package cluster

// postSet builds the set of node keys, so outlink filtering and edge
// construction only ever see targets that exist in the corpus.
func postSet(posts []post) map[string]bool {
	set := make(map[string]bool, len(posts))
	for _, p := range posts {
		set[PostKey(p.Section, p.Slug)] = true
	}
	return set
}
