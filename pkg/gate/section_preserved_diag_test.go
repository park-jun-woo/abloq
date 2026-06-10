//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what sectionPreservedDiag가 사라진 첫 섹션 키를 진단하고 유지 시 nil을 반환하는지 검증
package gate

import "testing"

func TestSectionPreservedDiag(t *testing.T) {
	a := &Article{Path: "p.md",
		Doc:  &Doc{Sections: []SectionHit{{Key: "related"}}},
		Base: &Doc{Sections: []SectionHit{{Key: "related"}, {Key: "sources", Text: "Sources"}}},
	}
	d := sectionPreservedDiag(a)
	if d == nil || d.Rule != "section-preserved" {
		t.Fatalf("diag = %v, want section-preserved", d)
	}
	kept := &Article{Path: "p.md",
		Doc:  &Doc{Sections: []SectionHit{{Key: "sources"}, {Key: "related"}}},
		Base: &Doc{Sections: []SectionHit{{Key: "related"}}},
	}
	if got := sectionPreservedDiag(kept); got != nil {
		t.Errorf("kept: want nil, got %v", got)
	}
}
