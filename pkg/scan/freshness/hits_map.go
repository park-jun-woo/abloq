//ff:func feature=scan type=parser control=iteration dimension=1
//ff:what 합계 목록 → 게이트 키(lang/section/slug) 인덱스 맵 — 스캐너의 우선순위 신호 조회용
package freshness

import "github.com/park-jun-woo/abloq/pkg/queueio"

// HitsMap indexes crawl-hit sums by the article join key so the scanner can
// look up the priority signal per entry.
func HitsMap(sums []HitSum) map[string]int64 {
	m := make(map[string]int64, len(sums))
	for _, s := range sums {
		m[queueio.JoinKey(s.Lang, s.Section, s.Slug)] = s.Hits
	}
	return m
}
