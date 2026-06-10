//ff:func feature=archive type=client control=iteration dimension=1
//ff:what GSC 제출 그룹 실행 — target마다 publish 호출, per-target 실패 격리
package archive

// gscSubmitAll publishes every target with the shared access token. Each
// target gets its own receipt — one failure never blocks the rest.
func gscSubmitAll(pending []Pending, endpoint, token string) []Item {
	items := make([]Item, 0, len(pending))
	for _, p := range pending {
		items = append(items, gscSubmit(p, endpoint, token))
	}
	return items
}
