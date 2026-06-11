//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what lessQueueItem이 priority 내림차순·동률 조인 키 오름차순으로 정렬하는지 검증
package common

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestLessQueueItem(t *testing.T) {
	hi := queueio.Item{Kind: "refresh", Lang: "en", Section: "posts", Slug: "z", Priority: 9}
	lo := queueio.Item{Kind: "refresh", Lang: "en", Section: "posts", Slug: "a", Priority: 1}
	if !lessQueueItem(hi, lo) || lessQueueItem(lo, hi) {
		t.Error("higher priority must sort first")
	}
	tieA := queueio.Item{Lang: "en", Section: "posts", Slug: "a", Priority: 5}
	tieB := queueio.Item{Lang: "en", Section: "posts", Slug: "b", Priority: 5}
	if !lessQueueItem(tieA, tieB) || lessQueueItem(tieB, tieA) {
		t.Error("ties must sort by join key ascending")
	}
}
