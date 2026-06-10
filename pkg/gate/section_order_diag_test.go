//ff:func feature=gate type=rule control=sequence
//ff:what sectionOrderDiag가 첫 역전 쌍의 위치·키를 진단에 담는지 검증
package gate

import "testing"

func TestSectionOrderDiag(t *testing.T) {
	rank := map[string]int{"related": 3, "sources": 5}
	a := &Article{Path: "p.md", Doc: &Doc{BodyStart: 5, Sections: []SectionHit{
		{Key: "sources", Text: "Sources", Line: 2},
		{Key: "related", Text: "Related", Line: 6},
	}}}
	d := sectionOrderDiag(rank, a)
	if d == nil {
		t.Fatal("want a diagnostic")
	}
	if d.Line != 11 || d.Rule != "section-order" {
		t.Errorf("diag = %+v, want line 11 rule section-order", d)
	}
	ordered := &Article{Path: "p.md", Doc: &Doc{Sections: []SectionHit{
		{Key: "related", Line: 2}, {Key: "sources", Line: 6},
	}}}
	if got := sectionOrderDiag(rank, ordered); got != nil {
		t.Errorf("ordered: want nil, got %v", got)
	}
}
