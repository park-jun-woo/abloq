//ff:func feature=quest type=rule control=sequence
//ff:what adaptRule 검증 — min-sources 발동이 Fact(파일:라인/기대/실제)로 매핑되고 클린 글은 발동하지 않는지
package writing

import (
	"strings"
	"testing"
)

func TestAdaptRule(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	r := adaptRule("min-sources")
	noSources := removeLine(removeLine(art, "## Sources"), "Internal style guide")
	fired, fact := fireRule(t, r, subWith(t, root, noSources, ""))
	if !fired {
		t.Fatal("min-sources: want fired on article without sources section")
	}
	if !strings.HasPrefix(fact.Where, "content/en/posts/fixture.md:") {
		t.Errorf("Where = %q", fact.Where)
	}
	if fact.Expected == "" || !strings.Contains(fact.Actual, "min_sources") {
		t.Errorf("Expected/Actual = %q / %q", fact.Expected, fact.Actual)
	}
	if fired, _ := fireRule(t, r, subWith(t, root, art, "")); fired {
		t.Error("clean article: min-sources fired")
	}
}
