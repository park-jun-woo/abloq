//ff:type feature=visibility type=scorer
//ff:what 우선순위 인터페이스 — 같은 신호면 같은 점수(결정적), Phase014 측정 기반 구현이 같은 계약으로 합류
package priority

// Scorer turns visibility signals into a queue priority. Implementations must
// be deterministic: identical signals yield an identical score — the queue
// file serialization and the CLI/endpoint equivalence both depend on it.
type Scorer interface {
	Score(s Signals) int64
}
