//ff:func feature=gate type=rule control=sequence
//ff:what [image-first] structure.order에 image 미선언 시 룰이 스킵되는지 검증
package gate

import "testing"

func TestRuleImageFirstSkip(t *testing.T) {
	b := loadGateBlog(t)
	b.Structure.Order = []string{"body", "sources"}
	a := artFromMD(t, b, "en", "tech", "no-image", "articles/no-image.md")
	tgt := NewTarget("testdata", b, []*Article{a})
	if diags := ruleImageFirst(tgt); len(diags) != 0 {
		t.Errorf("image undeclared: want 0 diagnostics, got %v", diags)
	}
}
