//ff:type feature=visibility type=schema topic=report
//ff:what 윈도 내 큐 적재 요약 1행 — kind·status별 건수, QueueItem.AggMonthCountsJson 행과 1:1
package report

// QueueCount is one (kind, status) queue tally over the report window
// (created_at, UTC dates).
type QueueCount struct {
	Kind   string `json:"kind"`
	Status string `json:"status"`
	Count  int64  `json:"count"`
}
