//ff:type feature=visibility type=schema topic=gsc
//ff:what GSC 일간 스냅샷 1행 — 일자·페이지·노출·클릭·평균순위, gsc_snapshots 1행과 1:1 (JSON 키 = 컬럼명)
//ff:why JSON 태그가 gsc_snapshots 컬럼명과 일치해야 한다 — 백엔드 @call 래퍼가 이 배열을 그대로 jsonb 업서트에 먹인다 (Phase013)
package gsc

// Snapshot is one (day, page) Search Analytics row. SnapDate is "YYYY-MM-DD"
// (UTC); JSON keys mirror the abloqd gsc_snapshots table columns.
type Snapshot struct {
	SnapDate    string  `json:"snap_date"`
	Page        string  `json:"page"`
	Impressions int64   `json:"impressions"`
	Clicks      int64   `json:"clicks"`
	AvgPosition float64 `json:"avg_position"`
}
