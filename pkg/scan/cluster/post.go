//ff:type feature=scan type=schema topic=cluster
//ff:what 그래프 노드 1건 — 기본 언어 글의 섹션/slug/날짜/태그와 해석된 아웃링크 키 목록
package cluster

// post is one default-language article as a graph node. Outlinks holds the
// resolved in-corpus targets as <section>/<slug> keys — deduplicated, in
// document order, self-links excluded — so in-degree and candidate exclusion
// read the same edge set.
type post struct {
	Section  string
	Slug     string
	Date     string
	Tags     []string
	Outlinks []string
}
