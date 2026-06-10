//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what 후보 정렬 비교자 — 태그 교집합 내림차순 → 동일 섹션 → 발행일 근접 → <section>/<slug> 사전순(최종 동률 키)
//ff:why slug만으로는 전순서가 아니다(섹션 간 slug 중복) — 최종 동률 키는 PostKey 사전순으로 고정해 CLI·endpoint 정렬이 갈라질 수 없게 한다
package cluster

// lessRanked orders candidates: more shared tags first, then same-section,
// then publication-date proximity, with the <section>/<slug> key as the
// total-order tie break.
func lessRanked(a, b ranked) bool {
	if a.cand.SharedTags != b.cand.SharedTags {
		return a.cand.SharedTags > b.cand.SharedTags
	}
	if a.sameSection != b.sameSection {
		return a.sameSection
	}
	if a.dateDist != b.dateDist {
		return a.dateDist < b.dateDist
	}
	return a.key < b.key
}
