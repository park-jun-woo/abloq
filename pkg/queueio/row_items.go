//ff:func feature=queueio type=generator control=iteration dimension=1
//ff:what 행 목록 → Item 목록 — 파일 기록 직전 DB id를 떼어낸다 (큐 파일에 id 불포함)
package queueio

// rowItems strips the database identity off rows before file serialization.
func rowItems(rows []Row) []Item {
	items := make([]Item, 0, len(rows))
	for _, r := range rows {
		items = append(items, r.Item)
	}
	return items
}
