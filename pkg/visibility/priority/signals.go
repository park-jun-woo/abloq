//ff:type feature=visibility type=schema
//ff:what 우선순위 산출 입력 신호 — 글의 발행일(front matter date 스칼라)과 crawl_hits 합계
package priority

// Signals carries the per-article visibility signals a Scorer consumes.
// Date is the front matter date scalar as-is (ISO-8601); Hits is the
// crawl_hits sum for the article (0 when no measurement data exists).
type Signals struct {
	Date string
	Hits int64
}
