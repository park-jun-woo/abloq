//ff:func feature=queueio type=generator control=sequence
//ff:what rowIDs가 id 목록을 순서대로 모으고 빈 입력에 nil 아닌 빈 슬라이스를 반환하는지 검증
package queueio

import "testing"

func TestRowIDs(t *testing.T) {
	ids := rowIDs([]Row{{ID: 5}, {ID: 7}})
	if len(ids) != 2 || ids[0] != 5 || ids[1] != 7 {
		t.Errorf("want [5 7], got %v", ids)
	}
	if got := rowIDs(nil); got == nil || len(got) != 0 {
		t.Errorf("empty input must yield non-nil empty slice, got %v", got)
	}
}
