//ff:func feature=visibility type=scorer control=sequence
//ff:what ColdStart 점수 — crawl_hits 합계가 있으면 그 값, 없으면 date의 epoch 일수(최신=큰 값=높은 우선순위)
package priority

// Score returns the crawl_hits sum when the article has any, otherwise the
// date recency score (days since the Unix epoch — newer dates score higher).
// The two scales are not comparable; that is acceptable for the cold-start
// fallback and is replaced by the Phase014 measurement-based Scorer.
func (ColdStart) Score(s Signals) int64 {
	if s.Hits > 0 {
		return s.Hits
	}
	return dateScore(s.Date)
}
