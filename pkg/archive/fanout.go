//ff:func feature=archive type=client control=iteration dimension=1
//ff:what 같은 배치 결과(response/status)를 영수증 항목들로 전개 — 일괄 API와 배치 단위 실패의 공통 환원
package archive

import "encoding/json"

// fanoutItems maps every pending receipt to one shared response/status —
// used by batch submissions (IndexNow) and batch-level failures (missing
// credentials, token failure, quota deferral).
func fanoutItems(pending []Pending, endpoint string, response json.RawMessage, status string) []Item {
	items := make([]Item, 0, len(pending))
	for _, p := range pending {
		items = append(items, Item{
			DeployID: p.DeployID,
			Kind:     p.Kind,
			Target:   p.Target,
			Request:  requestJSON(endpoint, p.Target),
			Response: response,
			Status:   status,
		})
	}
	return items
}
