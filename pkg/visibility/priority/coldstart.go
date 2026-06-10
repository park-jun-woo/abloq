//ff:type feature=visibility type=scorer
//ff:what 콜드스타트 스코어러 — 측정 데이터 無 상태의 영구 폴백 (임시물 아님), crawl_hits 합계 → 없으면 date 최신 우선
package priority

// ColdStart is the permanent fallback Scorer for articles (or whole
// instances) without measurement data: the crawl_hits sum when present,
// otherwise date recency — on a cold start, recency is the visibility proxy.
// Phase014 adds a measurement-based Scorer behind the same interface.
type ColdStart struct{}
