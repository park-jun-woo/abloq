//ff:func feature=quest type=rule control=sequence
//ff:what heading-canonical 어댑터 발동 검증 — "### Sources"(레벨 위반) 글에서 Fact 매핑
package writing

import (
	"strings"
	"testing"
)

func TestAdaptHeadingCanonical(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	bad := strings.Replace(art, "## Sources", "### Sources", 1)
	fired, fact := fireRule(t, adaptRule("heading-canonical"), subWith(t, root, bad, ""))
	if !fired {
		t.Fatal("heading-canonical: want fired on ### Sources")
	}
	if !strings.Contains(fact.Actual, "##") {
		t.Errorf("Actual = %q", fact.Actual)
	}
}
