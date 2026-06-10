//ff:type feature=visibility type=schema
//ff:what 우선순위 산출 입력 신호 — 발행일·crawl_hits 합계(콜드스타트)에 30d 측정 신호 4종(fetcher·train·GSC 노출·인용 적중) 확장 (Phase014)
package priority

// Signals carries the per-article visibility signals a Scorer consumes.
// Date is the front matter date scalar as-is (ISO-8601); Hits is the all-time
// crawl_hits sum (0 when no measurement data exists — the cold-start signal).
// The Phase014 measurement signals cover the 30-day window ending on the
// anchor month's last day: FetcherHits and TrainHits are the bot-category
// hit sums (search-category hits feed the report table only, never a
// Scorer), GSCTrend is the GSC impressions sum and CitationHits the
// cited-sample count of the same window.
type Signals struct {
	Date         string
	Hits         int64
	FetcherHits  int64
	TrainHits    int64
	GSCTrend     int64
	CitationHits int64
}
