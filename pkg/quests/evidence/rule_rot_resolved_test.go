//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what rot-resolved가 잔존 rot 인용을 FAIL로 잡고, 살아있는 출처로의 교체는 통과하는지 검증
package evidence

import (
	"strings"
	"testing"
)

func TestRuleRotResolved(t *testing.T) {
	r := ruleRotResolved()
	if r.Meta.ID != "rot-resolved" {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	// Empty diff: the rot URL is still cited — fires.
	root := writeInstance(t)
	fired, fact := r.Check(subWith(t, root))
	if !fired || !strings.Contains(fact.Actual, rotURL) {
		t.Errorf("untouched article: fired=%v fact=%+v", fired, fact)
	}
	// Replacing the citation with a live source resolves it.
	root2 := writeInstance(t)
	writeFile(t, root2, "content/en/posts/a.md", strings.Replace(baseArticleMD,
		rotURL, "https://example.org/live-study", 1))
	if fired, fact := r.Check(subWith(t, root2)); fired {
		t.Errorf("replaced citation fired: %+v", fact)
	}
}
