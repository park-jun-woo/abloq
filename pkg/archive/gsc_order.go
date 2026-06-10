//ff:func feature=archive type=client control=sequence
//ff:what GSC 제출 순서 결정 — 우선순위(신규 0 < 갱신 1 < 미인덱스 2) 안정 정렬, 입력은 보존
package archive

import "sort"

// gscOrder returns a copy of pending sorted new-posts-first (stable, so the
// receipts.id order is kept inside each priority class).
func gscOrder(pending []Pending) []Pending {
	ordered := append([]Pending(nil), pending...)
	sort.SliceStable(ordered, func(i, j int) bool {
		return gscPriority(ordered[i]) < gscPriority(ordered[j])
	})
	return ordered
}
