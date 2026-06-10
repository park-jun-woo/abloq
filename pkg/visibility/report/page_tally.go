//ff:type feature=visibility type=schema topic=report
//ff:what 글 1건의 GSC 집계 — 윈도 노출·클릭 합계 (델타 아님, 전월 대비는 리포트 표시 전용)
package report

// PageTally is one article's GSC window aggregate: the impressions and
// clicks sum of the window (not a delta — the month-over-month comparison
// is a report display concern only).
type PageTally struct {
	Impressions int64
	Clicks      int64
}
