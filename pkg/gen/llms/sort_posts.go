//ff:func feature=gen type=generator control=sequence
//ff:what 발행 글을 언어→섹션→날짜(내림차순) 고정 규칙으로 안정 정렬한 복사본 반환 — llms.txt 멱등성의 핵심
package llms

import "sort"

// sortPosts returns a sorted copy: blog.yaml language order, then section order,
// then newest date first, then slug. Input order never leaks into the output.
func sortPosts(posts []Post, langs, sections []string) []Post {
	langRank := rankOf(langs)
	sectionRank := rankOf(sections)
	sorted := make([]Post, len(posts))
	copy(sorted, posts)
	sort.SliceStable(sorted, func(i, j int) bool {
		return lessPost(sorted[i], sorted[j], langRank, sectionRank)
	})
	return sorted
}
