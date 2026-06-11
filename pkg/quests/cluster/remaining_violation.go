//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what 재스캔 산출에서 대상 글의 항목을 찾아 큐 지정 위반의 잔존 여부 판정 — 항목 없음=전부 해소 (cluster-resolved 전용)
package cluster

import "github.com/park-jun-woo/abloq/pkg/queueio"

// remainingViolation locates the fresh scan's item for the submission's
// article and reports the first queued violation kind it still carries.
// No item for the article means every violation resolved (other articles'
// items are irrelevant).
func remainingViolation(items []queueio.Item, sub *Submission) (string, bool) {
	for _, it := range items {
		if it.Lang != sub.Lang || it.Section != sub.Section || it.Slug != sub.Slug {
			continue
		}
		return matchViolation(it.Payload["violations"], sub.ViolRules)
	}
	return "", false
}
