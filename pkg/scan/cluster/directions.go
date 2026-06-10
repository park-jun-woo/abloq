//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what 위반 글-후보 쌍의 제안 가능 방향 산출 — 간선이 없는 쪽만 "out"/"in", 양방향 기존재면 빈 목록(후보 제외)
//ff:why 양방향 제안(Phase011): 고립 글은 나가는 링크와 들어올 글이 모두 필요하다 — 이미 있는 간선 방향은 제안하지 않는다
package cluster

// directions lists the proposed edges between the violating article and one
// candidate that do not exist yet: "out" (article → candidate) and "in"
// (candidate → article). An empty result means the pair is already fully
// connected and the candidate is dropped.
func directions(from, to post, edges map[string]bool) []string {
	fromKey := PostKey(from.Section, from.Slug)
	toKey := PostKey(to.Section, to.Slug)
	dirs := make([]string, 0, 2)
	if !edges[fromKey+"->"+toKey] {
		dirs = append(dirs, "out")
	}
	if !edges[toKey+"->"+fromKey] {
		dirs = append(dirs, "in")
	}
	return dirs
}
