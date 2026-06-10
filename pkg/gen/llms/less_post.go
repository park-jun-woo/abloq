//ff:func feature=gen type=generator control=sequence
//ff:what Post 정렬 비교자 — 언어 랭크 → 섹션 랭크 → 날짜 내림차순(최신 먼저) → slug 오름차순
package llms

// lessPost orders posts by language rank, section rank, newest date first, then slug.
func lessPost(a, b Post, langRank, sectionRank map[string]int) bool {
	if langRank[a.Lang] != langRank[b.Lang] {
		return langRank[a.Lang] < langRank[b.Lang]
	}
	if sectionRank[a.Section] != sectionRank[b.Section] {
		return sectionRank[a.Section] < sectionRank[b.Section]
	}
	if a.Date != b.Date {
		return a.Date > b.Date
	}
	return a.Slug < b.Slug
}
