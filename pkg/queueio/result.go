//ff:type feature=queueio type=schema
//ff:what export 1회전의 결과 — exported로 전이할 open 행 id들과 consumed로 전이할 exported 행 id들
package queueio

// Result reports one export cycle: ExportedIDs are the open rows whose files
// now live in the pushed clone (status → exported), ConsumedIDs are the
// exported rows whose files an agent deleted (status → consumed).
type Result struct {
	ExportedIDs []int64
	ConsumedIDs []int64
}
