//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what isolationViolation이 인링크 0 글만 적발하는지 검증
package cluster

import "testing"

func TestIsolationViolation(t *testing.T) {
	v := isolationViolation(0)
	if v == nil || v.Rule != "no-isolated-post" || v.Detail != "no inbound internal links" {
		t.Errorf("violation = %+v", v)
	}
	if v := isolationViolation(1); v != nil {
		t.Errorf("linked article flagged: %+v", v)
	}
}
