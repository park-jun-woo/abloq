//ff:func feature=archive type=client control=iteration dimension=1
//ff:what 변경 URL × kind 3종을 status=pending 영수증 항목으로 전개
package archive

import "encoding/json"

// planItems fans the changed URLs of one deploy out to every receipt kind.
// Request/response stay empty at the pending stage — the processor records
// the actual exchange.
func planItems(deployID string, changed []string) []Item {
	items := make([]Item, 0, len(changed)*len(Kinds))
	for _, target := range changed {
		for _, kind := range Kinds {
			items = append(items, Item{
				DeployID: deployID,
				Kind:     kind,
				Target:   target,
				Request:  json.RawMessage(`{}`),
				Response: json.RawMessage(`{}`),
				Status:   StatusPending,
			})
		}
	}
	return items
}
