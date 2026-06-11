//ff:func feature=quest type=rule control=sequence
//ff:what front-matter-intact 어댑터 발동 검증 — Base 부착 픽스처에서 title 변경(lastmod 외 변경) 시 Fact 매핑
package writing

import (
	"strings"
	"testing"
)

func TestAdaptFrontMatterIntact(t *testing.T) {
	root := writeInstance(t)
	base, _ := passFixtures()
	cur := strings.Replace(base, `title: "Test Article"`, `title: "Renamed"`, 1)
	fired, fact := fireRule(t, adaptRule("front-matter-intact"), subWith(t, root, cur, base))
	if !fired {
		t.Fatal("front-matter-intact: want fired when title changes vs the baseline")
	}
	if !strings.Contains(fact.Actual, "front matter") {
		t.Errorf("Actual = %q", fact.Actual)
	}
}
