//ff:func feature=gate type=rule control=sequence
//ff:what presentKeys가 섹션 헤딩 목록을 키 집합으로 만드는지 검증
package gate

import "testing"

func TestPresentKeys(t *testing.T) {
	got := presentKeys([]SectionHit{{Key: "sources"}, {Key: "related"}, {Key: "sources"}})
	if len(got) != 2 || !got["sources"] || !got["related"] {
		t.Errorf("presentKeys = %v", got)
	}
	if got := presentKeys(nil); len(got) != 0 {
		t.Errorf("empty: want empty set, got %v", got)
	}
}
