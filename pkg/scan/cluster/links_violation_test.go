//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what linksViolation이 임계 미달 아웃링크만 적발하는지(경계 포함) 검증
package cluster

import "testing"

func TestLinksViolation(t *testing.T) {
	v := linksViolation(1, 2)
	if v == nil || v.Rule != "min-internal-links" || v.Detail != "outbound internal links 1 below min 2" {
		t.Errorf("violation = %+v", v)
	}
	if v := linksViolation(2, 2); v != nil {
		t.Errorf("boundary flagged: %+v", v)
	}
	if v := linksViolation(0, 0); v != nil {
		t.Errorf("zero threshold flagged: %+v", v)
	}
}
