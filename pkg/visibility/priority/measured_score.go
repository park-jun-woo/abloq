//ff:func feature=visibility type=scorer control=sequence
//ff:what Measured 점수 — w_fetcher·fetcher + w_train·train + w_gsc·노출 + w_cite·인용 적중의 정수 가중 합 (순수·결정적)
package priority

// Score returns the weighted measurement sum. Pure integer arithmetic over
// the signals — identical signals always yield an identical score.
func (m Measured) Score(s Signals) int64 {
	return m.W.Fetcher*s.FetcherHits +
		m.W.Train*s.TrainHits +
		m.W.GSC*s.GSCTrend +
		m.W.Citation*s.CitationHits
}
