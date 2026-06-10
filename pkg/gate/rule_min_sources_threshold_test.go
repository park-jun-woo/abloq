//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what [min-sources] 임계·스킵 케이스 — min_sources 2 PASS / 3 FAIL / 0 스킵, structure에 sources 미선언이면 스킵
package gate

import "testing"

func TestRuleMinSourcesThreshold(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "fixture", "evidence/sources-ok.md")
	b.Geo.MinSources = 2
	checkDiags(t, ruleMinSources(NewTarget("testdata", b, []*Article{a})), 0, "", "")
	b.Geo.MinSources = 3
	checkDiags(t, ruleMinSources(NewTarget("testdata", b, []*Article{a})), 1, "min-sources", "requires >= 3")
	b.Geo.MinSources = 0
	checkDiags(t, ruleMinSources(NewTarget("testdata", b, []*Article{a})), 0, "", "")
	b.Geo.MinSources = 1
	b.Structure.Order = []string{"body"}
	checkDiags(t, ruleMinSources(NewTarget("testdata", b, []*Article{a})), 0, "", "")
}
