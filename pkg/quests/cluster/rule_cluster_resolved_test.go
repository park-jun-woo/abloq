//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what cluster-resolved가 무작업(빈 diff)을 FAIL로 잡고, out+in 링크 추가로 위반이 소멸하면 통과하는지 검증 (부분 해소는 잔존 FAIL)
package cluster

import (
	"strings"
	"testing"
)

func TestRuleClusterResolved(t *testing.T) {
	r := ruleClusterResolved()
	if r.Meta.ID != "cluster-resolved" {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	// Empty diff: both violations remain — fires.
	root := writeInstance(t)
	fired, fact := r.Check(subWith(t, root))
	if !fired || !strings.Contains(fact.Actual, "min-internal-links") {
		t.Errorf("untouched graph: fired=%v fact=%+v", fired, fact)
	}
	// Outbound link only: isolation remains — still fires.
	root2 := writeInstance(t)
	writeFile(t, root2, "content/en/posts/thin.md",
		thinArticleMD+"\nSee the [hub](/posts/hub/) overview.\n")
	fired, fact = r.Check(subWith(t, root2))
	if !fired || !strings.Contains(fact.Actual, "no-isolated-post") {
		t.Errorf("partial fix: fired=%v fact=%+v", fired, fact)
	}
	// Outbound + inbound (candidate-side anchor): both resolve.
	root3 := writeInstance(t)
	writeFile(t, root3, "content/en/posts/thin.md",
		thinArticleMD+"\nSee the [hub](/posts/hub/) overview.\n")
	writeFile(t, root3, "content/en/posts/hub.md",
		hubArticleMD+"\nThe [thin](/posts/thin/) article covers the edge case.\n")
	if fired, fact := r.Check(subWith(t, root3)); fired {
		t.Errorf("resolved graph fired: %+v", fact)
	}
}
