//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what [section-preserved] 원본 없음(nil)과 미변경(공유 Doc) 글이 스킵되는지 검증
package gate

import "testing"

func TestRuleSectionPreservedSkip(t *testing.T) {
	b := loadGateBlog(t)
	dropped := artFromMD(t, b, "en", "tech", "base", "baseline/dropped-section.md")
	tgt := NewTarget("testdata", b, []*Article{dropped})
	if diags := ruleSectionPreserved(tgt); len(diags) != 0 {
		t.Errorf("nil baseline: want skip, got %v", diags)
	}
	dropped.Base = dropped.Doc
	if diags := ruleSectionPreserved(tgt); len(diags) != 0 {
		t.Errorf("shared baseline: want skip, got %v", diags)
	}
}
