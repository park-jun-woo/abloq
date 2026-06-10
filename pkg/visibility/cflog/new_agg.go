//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 빈 수집 누적기 생성 — URI 역매핑을 받아 카운터 맵 초기화
package cflog

// NewAgg builds an empty aggregation accumulator over the given URI reverse
// map.
func NewAgg(urls map[string]Article) *Agg {
	return &Agg{
		URLs:    urls,
		hits:    map[hitKey]*hitCount{},
		unknown: map[string]*unknownAgg{},
		Raw:     map[string]int64{},
	}
}
