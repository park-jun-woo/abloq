//ff:func feature=visibility type=scorer control=sequence
//ff:what Composite 점수 — 측정 신호 4종 전부 0이면 ColdStart 폴백(무변형), 아니면 Measured 가중 합
package priority

// Score routes between the two scorers. The fallback condition is "every
// measurement signal is zero": an instance (or article) without measurement
// data behaves exactly like the Phase009 cold start.
func (c Composite) Score(s Signals) int64 {
	if s.FetcherHits == 0 && s.TrainHits == 0 && s.GSCTrend == 0 && s.CitationHits == 0 {
		return ColdStart{}.Score(s)
	}
	return Measured{W: c.W}.Score(s)
}
