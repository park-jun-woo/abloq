//ff:type feature=queueio type=schema
//ff:what queue_items 행 1건 — DB id + Item, exporter가 상태 전이(open→exported, exported→consumed) 대상 식별에 사용
package queueio

// Row is a queue item joined with its database identity. The ID never reaches
// a queue file — it only routes the status transitions back to the table.
type Row struct {
	ID int64 `json:"id"`
	Item
}
