//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 파싱된 큐 항목 1개 → quest Item — 대상 글 실재 확인, keys 부재(legacy) 시 선언 언어로 보충, payload 고정 (SeedQueue 전용)
package common

import (
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// seedQueueItem turns one parsed queue item into a quest TODO item. The
// target article must exist (queue items point at committed articles); a
// queue file without a keys: block (legacy export) falls back to the
// declared languages so the queue-scope allowed set never shrinks below the
// translation companions.
func seedQueueItem(root string, qit queueio.Item, langs []string) (*quest.Item, error) {
	article, err := ArticlePath(root, qit.Lang, qit.Section, qit.Slug)
	if err != nil {
		return nil, err
	}
	keys := qit.Keys
	if len(keys) == 0 {
		keys = queueio.KeysFor(langs, qit.Section, qit.Slug)
	}
	it := &quest.Item{Key: queueio.JoinKey(qit.Lang, qit.Section, qit.Slug), State: quest.TODO}
	p := QueuePayload{Root: root, Article: article, Lang: qit.Lang,
		Section: qit.Section, Slug: qit.Slug, Keys: keys, Queue: qit.Payload}
	if err := it.SetPayload(p); err != nil {
		return nil, err
	}
	return it, nil
}
