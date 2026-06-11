//ff:func feature=quest type=rule control=sequence
//ff:what AdaptRule 검증 — TargetCarrier 제출물 위에서 min-sources 발동이 Fact로 매핑되고 클린 글은 발동하지 않는지
package common

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestAdaptRule(t *testing.T) {
	root, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	noSources := strings.ReplaceAll(fixtureArticleMD, "## Sources", "## Notes")
	bad, _ := writeFixture(t, "content/en/posts/a.md", noSources)
	r := AdaptRule("min-sources")
	if r.Meta.ID != "min-sources" || r.Meta.Level != rgate.LevelFail || r.Meta.Desc == "" {
		t.Errorf("Meta = %+v", r.Meta)
	}
	tgt, _, err := AssembleTarget(bad, "content/en/posts/a.md", "en", "posts", "a")
	if err != nil {
		t.Fatalf("AssembleTarget: %v", err)
	}
	fired, fact := r.Check(rgate.Context{Submission: &carrier{tgt}})
	if !fired || !strings.HasPrefix(fact.Where, "content/en/posts/a.md:") {
		t.Errorf("fired=%v fact=%+v — want fired with located Fact", fired, fact)
	}
	clean, _, err := AssembleTarget(root, "content/en/posts/a.md", "en", "posts", "a")
	if err != nil {
		t.Fatalf("AssembleTarget clean: %v", err)
	}
	if fired, _ := r.Check(rgate.Context{Submission: &carrier{clean}}); fired {
		t.Error("clean article: min-sources fired")
	}
}
