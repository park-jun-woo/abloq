//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what 게이트 FAIL 경로 검증 — 빈 diff(무작업)는 cluster-resolved가, candidates 밖 글 수정은 queue-scope가 잡는지
package cluster

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluateFail(t *testing.T) {
	// Empty diff — no work done: the resolution re-check is the root cause.
	root := writeInstance(t)
	v := rgate.Evaluate(Definition{}.Rules(), subWith(t, root))
	if v.Outcome != quest.OutFail || v.RootCause != "cluster-resolved" {
		t.Fatalf("empty diff: outcome=%s root=%s — want FAIL/cluster-resolved", v.Outcome, v.RootCause)
	}
	// A resolution that ALSO edits a non-candidate article: queue-scope fires.
	root2 := writeInstance(t)
	writeFile(t, root2, "content/en/posts/thin.md",
		thinArticleMD+"\nSee the [hub](/posts/hub/) overview.\n")
	writeFile(t, root2, "content/en/posts/hub.md",
		hubArticleMD+"\nThe [thin](/posts/thin/) article covers the edge case.\n")
	writeFile(t, root2, "content/en/posts/extra.md", extraArticleMD+"\nStray edit.\n")
	v = rgate.Evaluate(Definition{}.Rules(), subWith(t, root2))
	if v.Outcome != quest.OutFail {
		t.Fatalf("non-candidate edit: outcome=%s — want FAIL", v.Outcome)
	}
	found := false
	for _, f := range v.Facts {
		if f.Rule == "queue-scope" && strings.Contains(f.Actual, "content/en/posts/extra.md") {
			found = true
		}
	}
	if !found {
		t.Errorf("queue-scope Fact missing: %+v", v.Facts)
	}
}
