//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what [front-matter-intact] 현재본의 front matter 블록 소실을 진단하는지 검증
package gate

import "testing"

func TestRuleFrontMatterIntactMissing(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "base", "articles/no-fm.md")
	a.Base = artFromMD(t, b, "en", "tech", "base", "baseline/base.md").Doc
	tgt := NewTarget("testdata", b, []*Article{a})
	checkDiags(t, ruleFrontMatterIntact(tgt), 1, "front-matter-intact", "missing or malformed")
}
