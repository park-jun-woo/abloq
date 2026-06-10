//ff:func feature=gate type=parser control=iteration dimension=1 topic=lossless
//ff:what structuralLines가 이미지/저작자 표기/정상·비정상 레벨 섹션 헤딩 라인 인덱스를 모두 수집하는지 검증
package gate

import "testing"

func TestStructuralLines(t *testing.T) {
	d := &Doc{
		FirstContentLine: 1, FirstIsImage: true, AttribLine: 2,
		Sections: []SectionHit{{Key: "sources", Line: 5}},
		BadLevel: []SectionHit{{Key: "related", Line: 8}},
	}
	skip := structuralLines(d)
	for _, idx := range []int{1, 2, 5, 8} {
		if !skip[idx] {
			t.Errorf("want line %d skipped", idx)
		}
	}
	if len(skip) != 4 {
		t.Errorf("want exactly 4 skipped lines, got %v", skip)
	}
	if got := structuralLines(&Doc{FirstContentLine: 0, AttribLine: -1}); len(got) != 0 {
		t.Errorf("plain doc: want no skips, got %v", got)
	}
}
