//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what lastmod-advance가 미전진·누락 lastmod를 FAIL로 잡고 전진은 통과하는지 검증 (빈 diff 무작업 통과 차단)
package refresh

import (
	"strings"
	"testing"
)

func TestRuleLastmodAdvance(t *testing.T) {
	r := ruleLastmodAdvance()
	if r.Meta.ID != "lastmod-advance" {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	// Empty diff: lastmod unchanged — the no-work cheese fires.
	root := writeInstance(t)
	fired, fact := r.Check(subWith(t, root))
	if !fired || !strings.Contains(fact.Actual, "did not advance") {
		t.Errorf("unchanged lastmod: fired=%v fact=%+v", fired, fact)
	}
	// Advanced lastmod passes.
	root2 := writeInstance(t)
	writeFile(t, root2, "content/en/posts/a.md",
		strings.Replace(baseArticleMD, "lastmod: 2026-06-02", "lastmod: 2026-06-09", 1))
	if fired, fact := r.Check(subWith(t, root2)); fired {
		t.Errorf("advanced lastmod fired: %+v", fact)
	}
	// Missing lastmod fails.
	root3 := writeInstance(t)
	writeFile(t, root3, "content/en/posts/a.md",
		strings.Replace(baseArticleMD, "lastmod: 2026-06-02\n", "", 1))
	if fired, fact := r.Check(subWith(t, root3)); !fired || !strings.Contains(fact.Actual, "missing") {
		t.Errorf("missing lastmod: fired=%v fact=%+v", fired, fact)
	}
}
