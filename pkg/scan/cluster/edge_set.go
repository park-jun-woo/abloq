//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what 방향 간선 집합 조립 — "<fromKey>-><toKey>" 키, 후보 제안의 "이미 링크된 글 제외" 판정 입력
package cluster

// edgeSet builds the directed edge set over the finalized outlinks. The
// candidate ranking consults it to drop directions that already exist.
func edgeSet(posts []post) map[string]bool {
	edges := map[string]bool{}
	for _, p := range posts {
		from := PostKey(p.Section, p.Slug)
		for _, to := range p.Outlinks {
			edges[from+"->"+to] = true
		}
	}
	return edges
}
