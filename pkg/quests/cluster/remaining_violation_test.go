//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what remainingViolation이 대상 글 항목만 대조하고(타 글 무관) 항목 부재는 전부 해소로 보는지 검증
package cluster

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestRemainingViolation(t *testing.T) {
	sub := &Submission{Consumption: &common.Consumption{},
		Lang: "en", Section: "posts", Slug: "thin",
		ViolRules: map[string]bool{"no-isolated-post": true}}
	other := queueio.Item{Kind: "cluster", Lang: "en", Section: "posts", Slug: "other",
		Payload: map[string]string{"violations": `[{"rule":"no-isolated-post","detail":"d"}]`}}
	if rule, remains := remainingViolation([]queueio.Item{other}, sub); remains {
		t.Errorf("another article's violation matched: %s", rule)
	}
	mine := queueio.Item{Kind: "cluster", Lang: "en", Section: "posts", Slug: "thin",
		Payload: map[string]string{"violations": `[{"rule":"no-isolated-post","detail":"d"}]`}}
	if rule, remains := remainingViolation([]queueio.Item{other, mine}, sub); !remains || rule != "no-isolated-post" {
		t.Errorf("own violation missed: %s %v", rule, remains)
	}
	if _, remains := remainingViolation(nil, sub); remains {
		t.Error("no item for the article means resolved")
	}
}
