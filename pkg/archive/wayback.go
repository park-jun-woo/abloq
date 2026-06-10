//ff:func feature=archive type=client control=iteration dimension=1
//ff:what Wayback 그룹 처리 — target마다 SPN2 저장을 실행, per-target 실패 격리
package archive

// processWayback saves every target via the SPN2 API. Each target gets its
// own receipt — one failure never blocks the rest.
func processWayback(pending []Pending) []Item {
	items := make([]Item, 0, len(pending))
	for _, p := range pending {
		items = append(items, saveWayback(p))
	}
	return items
}
