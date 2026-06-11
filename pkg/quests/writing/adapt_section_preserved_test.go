//ff:func feature=quest type=rule control=sequence
//ff:what section-preserved 어댑터 발동 검증 — Base 부착 픽스처에서 기준선 섹션(changelog) 삭제 시 Fact 매핑
package writing

import "testing"

func TestAdaptSectionPreserved(t *testing.T) {
	root := writeInstance(t)
	base, _ := passFixtures()
	cur := removeLine(removeLine(base, "## Changelog"), "2026-06-02 first version")
	fired, fact := fireRule(t, adaptRule("section-preserved"), subWith(t, root, cur, base))
	if !fired {
		t.Fatal("section-preserved: want fired when a baseline section is dropped")
	}
	if fact.Where == "" || fact.Actual == "" {
		t.Errorf("Fact incomplete: %+v", fact)
	}
}
