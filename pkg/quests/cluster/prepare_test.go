//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Prepare가 payload 위반 룰 집합을 적재하고 candidates 글 경로로 허용 집합을 확장하는지 검증 (violations 부재는 에러)
package cluster

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepare(t *testing.T) {
	root := writeInstance(t)
	ctx := subWith(t, root)
	sub := ctx.Submission.(*Submission)
	if !sub.ViolRules["min-internal-links"] || !sub.ViolRules["no-isolated-post"] || len(sub.ViolRules) != 2 {
		t.Errorf("viol rules = %v", sub.ViolRules)
	}
	if !sub.Allowed["content/en/posts/thin.md"] || !sub.Allowed["content/en/posts/hub.md"] {
		t.Errorf("allowed set must include the target and the candidates: %v", sub.Allowed)
	}
	if sub.Allowed["content/en/posts/extra.md"] {
		t.Error("a non-candidate article must stay out of the allowed set")
	}
	// A queue payload without violations is tampering or an issuance bug.
	it := &quest.Item{Key: "en/posts/thin", State: quest.TODO}
	p := common.QueuePayload{Root: root, Article: "content/en/posts/thin.md",
		Lang: "en", Section: "posts", Slug: "thin", Keys: []string{"en/posts/thin"},
		Queue: map[string]string{}}
	if err := it.SetPayload(p); err != nil {
		t.Fatal(err)
	}
	if _, _, err := (Definition{}).Prepare(nil, it, []byte(`{"article":"content/en/posts/thin.md"}`)); err == nil {
		t.Error("payload without violations: want error")
	}
}
