//ff:type feature=visibility type=schema topic=report
//ff:what 리포트 글별 표 1행 — 분류별 크롤·md·GSC 노출·클릭·cited·우선순위 점수, JSON 산출과 1:1
package report

// Row is one article line of the report table. Priority is the Composite
// scorer's value over the same window signals the row displays — the report
// is the scanner's priority input made visible.
type Row struct {
	Lang        string `json:"lang"`
	Section     string `json:"section"`
	Slug        string `json:"slug"`
	Date        string `json:"date"`
	Training    int64  `json:"training_hits"`
	Search      int64  `json:"search_hits"`
	Fetch       int64  `json:"fetch_hits"`
	MDHits      int64  `json:"md_hits"`
	Impressions int64  `json:"gsc_impressions"`
	Clicks      int64  `json:"gsc_clicks"`
	Cited       int64  `json:"cited_samples"`
	Priority    int64  `json:"priority"`
}
