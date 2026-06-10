//ff:type feature=visibility type=schema topic=gsc
//ff:what GSC 수집 1회 결과 — 스냅샷 행 묶음과 수집한 닫힌 일자 수
package gsc

// Result is one incremental GSC collection: every (day, page) row plus the
// number of closed days the run covered. Days == 0 means the cursor was
// already at the last closed day — the idempotency oracle.
type Result struct {
	Rows []Snapshot
	Days int64
}
