//ff:type feature=visibility type=schema topic=report
//ff:what 윈도 전체 합계 — 크롤 히트(분류 합)·md·GSC 노출·클릭·cited, 전월 대비 추이의 비교 단위
package report

// Totals is one window's whole-blog aggregate: the month-over-month trend
// compares the current and previous window's Totals.
type Totals struct {
	CrawlHits   int64 `json:"crawl_hits"`
	MDHits      int64 `json:"md_hits"`
	Impressions int64 `json:"gsc_impressions"`
	Clicks      int64 `json:"gsc_clicks"`
	Cited       int64 `json:"cited_samples"`
}
