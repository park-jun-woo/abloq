//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what 글 1편의 아웃링크 확정 — 코퍼스에 실재하는 키만, 자기 참조 제외
//ff:why 삭제된 글로 향한 링크는 클러스터 연결이 아니다 — 그래프 차수(min-internal-links·no-isolated-post)는 실재 간선만 센다
package cluster

// filterOutlinks finalizes one node's outlinks: only targets that exist in
// the corpus count as cluster edges, and a self-link is no connection.
func filterOutlinks(p post, set map[string]bool) []string {
	self := PostKey(p.Section, p.Slug)
	out := make([]string, 0, len(p.Outlinks))
	for _, key := range p.Outlinks {
		if key == self || !set[key] {
			continue
		}
		out = append(out, key)
	}
	return out
}
