//ff:func feature=quest type=rule control=sequence
//ff:what front-matter-schema 어댑터 발동 검증 — tags 누락 글에서 Fact 매핑
package writing

import (
	"strings"
	"testing"
)

func TestAdaptFrontMatterSchema(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	noTags := removeLine(art, "tags: [test]")
	fired, fact := fireRule(t, adaptRule("front-matter-schema"), subWith(t, root, noTags, ""))
	if !fired {
		t.Fatal("front-matter-schema: want fired without tags")
	}
	if !strings.Contains(fact.Actual, "tags") {
		t.Errorf("Actual = %q", fact.Actual)
	}
}
