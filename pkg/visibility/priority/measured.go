//ff:type feature=visibility type=scorer
//ff:what 측정 기반 스코어러 — 30d 측정 신호 4종의 가중 합, 계수는 geo.priority_weights (Phase014)
package priority

// Measured scores an article from its measurement signals: the weighted sum
// of fetcher hits, train hits, GSC impressions and citation-sample hits over
// the 30-day window. Search-category hits are intentionally absent (§6.1
// names train bots and fetchers as the signals; search rides in the report
// table only). Citation samples feed this score and nothing else — §6.3
// allows exactly this single use, never a gate or verdict.
type Measured struct {
	W Weights
}
