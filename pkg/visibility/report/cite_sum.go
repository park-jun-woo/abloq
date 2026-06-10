//ff:type feature=visibility type=schema topic=report
//ff:what 윈도 내 글별 인용 샘플 합계 1행 — cited 적중 수와 전체 수, CitationSample.AggMonthJson 행과 1:1
package report

// CiteSum is one per-article citation-sample sum over the report window:
// Cited counts the cited=true samples, Total every sample of the article's
// queries. Trend record + priority input only — never a gate (§6.3).
type CiteSum struct {
	Lang    string `json:"lang"`
	Section string `json:"section"`
	Slug    string `json:"slug"`
	Cited   int64  `json:"cited"`
	Total   int64  `json:"total"`
}
