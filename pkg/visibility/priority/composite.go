//ff:type feature=visibility type=scorer
//ff:what 합성 스코어러 — 측정 신호가 하나라도 있으면 Measured, 전부 0이면 ColdStart 점수를 변형 없이 반환 (Phase014)
package priority

// Composite is the production Scorer: the measurement-based score when any
// of the four measurement signals is non-zero, the ColdStart score verbatim
// otherwise. No scale conversion happens on the fallback path — the
// cold-start date-recency values (epoch days) pass through untouched, which
// the scenario-freshness priority asserts (20607/20605) depend on.
type Composite struct {
	W Weights
}
