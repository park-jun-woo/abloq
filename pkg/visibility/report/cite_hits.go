//ff:func feature=visibility type=parser control=iteration dimension=1 topic=report
//ff:what 인용 합계 → 게이트 키 인덱스의 cited 적중 수 맵 — 우선순위 신호용 (§6.3 허용 단일 용도)
package report

import "github.com/park-jun-woo/abloq/pkg/queueio"

// CiteHits indexes the cited-sample counts by the article join key — the
// citation signal of the priority scorer, §6.3's single allowed use.
func CiteHits(sums []CiteSum) map[string]int64 {
	m := make(map[string]int64, len(sums))
	for _, s := range sums {
		m[queueio.JoinKey(s.Lang, s.Section, s.Slug)] += s.Cited
	}
	return m
}
