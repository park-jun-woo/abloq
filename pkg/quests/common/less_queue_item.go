//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 큐 항목 정렬 비교기 — priority 내림차순, 동률은 조인 키 오름차순 (결정적 Seed 순서)
package common

import "github.com/park-jun-woo/abloq/pkg/queueio"

// lessQueueItem orders queue items for seeding: highest priority first, join
// key ascending on ties so the seed order is deterministic for a given queue.
func lessQueueItem(a, b queueio.Item) bool {
	if a.Priority != b.Priority {
		return a.Priority > b.Priority
	}
	return queueio.JoinKey(a.Lang, a.Section, a.Slug) < queueio.JoinKey(b.Lang, b.Section, b.Slug)
}
