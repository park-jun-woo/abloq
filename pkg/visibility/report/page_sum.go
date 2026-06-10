//ff:type feature=visibility type=schema topic=report
//ff:what 윈도 내 GSC page별 노출·클릭 합계 1행 — GscSnapshot.AggPageMonthJson 행과 1:1, 글 귀속은 URL맵이 한다 (SQL 매핑 금지)
package report

// PageSum is one per-page GSC sum over the report window, as the backend's
// GscSnapshot.AggPageMonthJson emits it. Page is the full URL exactly as
// gsc_snapshots stores it; the page→article attribution happens in Go via
// the repository URL map (PageTotals) — URL composition rules live only in
// blog.yaml, so SQL must never join on them.
type PageSum struct {
	Page        string `json:"page"`
	Impressions int64  `json:"impressions"`
	Clicks      int64  `json:"clicks"`
}
