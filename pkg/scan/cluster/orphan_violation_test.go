//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what orphanViolation이 보유 글 1편짜리 태그만 적발하는지 검증
package cluster

import "testing"

func TestOrphanViolation(t *testing.T) {
	counts := map[string]int64{"geo": 3, "lonely": 1, "solo": 1}
	v := orphanViolation([]string{"geo", "lonely", "solo"}, counts)
	if v == nil || v.Rule != "no-orphan-tag" || v.Detail != "tags used by this article only: lonely, solo" {
		t.Errorf("violation = %+v", v)
	}
	if v := orphanViolation([]string{"geo"}, counts); v != nil {
		t.Errorf("shared tag flagged: %+v", v)
	}
	if v := orphanViolation(nil, counts); v != nil {
		t.Errorf("tagless article flagged: %+v", v)
	}
}
