//ff:func feature=gate type=rule control=sequence
//ff:what orderHas가 structure.order 선언 여부를 판정하는지 검증
package gate

import "testing"

func TestOrderHas(t *testing.T) {
	b := loadGateBlog(t)
	if !orderHas(b, "image") {
		t.Error("want image declared")
	}
	if orderHas(b, "toc") {
		t.Error("want toc undeclared")
	}
	b.Structure.Order = nil
	if orderHas(b, "image") {
		t.Error("empty order: want false")
	}
}
