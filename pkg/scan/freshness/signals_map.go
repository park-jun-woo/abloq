//ff:func feature=scan type=parser control=iteration dimension=1
//ff:what 합계 목록 → 게이트 키(lang/section/slug) 인덱스의 신호 맵 — Hits(전기간 합계, 콜드스타트 신호)만 채운 기본 신호
package freshness

import (
	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// SignalsMap indexes crawl-hit sums by the article join key as the base
// signals map: only Hits (the all-time sum, the cold-start signal) is
// filled. The backend merges the Phase014 measurement signals on top
// (report.MergeSignals); the CLI passes the result through unchanged.
func SignalsMap(sums []HitSum) map[string]priority.Signals {
	m := make(map[string]priority.Signals, len(sums))
	for _, s := range sums {
		m[queueio.JoinKey(s.Lang, s.Section, s.Slug)] = priority.Signals{Hits: s.Hits}
	}
	return m
}
