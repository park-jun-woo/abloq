//ff:func feature=visibility type=parser control=iteration dimension=1 topic=report
//ff:what 봇별 합계 → 게이트 키 인덱스의 분류별 집계 맵 — pkg/bots.CategoryOf로 분류, 사전 밖 봇은 버린다
package report

import (
	"github.com/park-jun-woo/abloq/pkg/bots"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// BotTotals classifies the per-bot window sums through the pkg/bots
// dictionary and indexes the per-article tallies by the article join key.
// A bot name outside the dictionary is dropped (crawl_hits only ever stores
// dictionary bots, so this is a defensive guard, not a filter).
func BotTotals(sums []BotSum) map[string]Tally {
	m := make(map[string]Tally, len(sums))
	for _, s := range sums {
		category, ok := bots.CategoryOf(s.Bot)
		if !ok {
			continue
		}
		key := queueio.JoinKey(s.Lang, s.Section, s.Slug)
		t := m[key]
		t.add(category, s.Hits, s.MDHits)
		m[key] = t
	}
	return m
}
