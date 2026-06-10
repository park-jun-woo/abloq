//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what 그래프 노드 키·최종 동률 정렬 키 조립 — <section>/<slug> (slug만으로는 전순서가 아님)
//ff:why CLI와 endpoint가 같은 pkg의 이 함수로 동률을 끊어야 diff -r 0 등가가 성립한다 — 호출부마다 키를 재조립하면 정렬이 갈라질 수 있다
package cluster

// PostKey builds the graph node key and the final sort tie-break key for one
// article: <section>/<slug>. Slug alone is not a total order — two sections
// may carry the same slug.
func PostKey(section, slug string) string {
	return section + "/" + slug
}
