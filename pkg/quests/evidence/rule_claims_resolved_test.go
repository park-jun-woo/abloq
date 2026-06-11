//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what claims-resolved가 무작업(빈 diff)을 FAIL로 잡고, 큐 주장에 출처를 단 제출은 통과하는지 검증
package evidence

import (
	"strings"
	"testing"
)

func TestRuleClaimsResolved(t *testing.T) {
	r := ruleClaimsResolved()
	if r.Meta.ID != "claims-resolved" {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	// Empty diff: the queued claim is still unsourced — fires.
	root := writeInstance(t)
	fired, fact := r.Check(subWith(t, root))
	if !fired || !strings.Contains(fact.Actual, "40%") {
		t.Errorf("untouched article: fired=%v fact=%+v", fired, fact)
	}
	// Sourcing the queued claim resolves it.
	root2 := writeInstance(t)
	writeFile(t, root2, "content/en/posts/a.md", strings.Replace(baseArticleMD,
		unsourcedClaim,
		unsourcedClaim+" [Migration report](https://example.org/spec)", 1))
	if fired, fact := r.Check(subWith(t, root2)); fired {
		t.Errorf("sourced claim fired: %+v", fact)
	}
}
