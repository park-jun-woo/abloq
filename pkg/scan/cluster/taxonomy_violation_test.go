//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what taxonomyViolation이 taxonomy 밖 태그만 적발하고 미선언 taxonomy는 스킵하는지 검증
package cluster

import "testing"

func TestTaxonomyViolation(t *testing.T) {
	taxonomy := []string{"geo", "abloq"}
	if v := taxonomyViolation([]string{"geo", "abloq"}, taxonomy); v != nil {
		t.Errorf("in-taxonomy tags flagged: %+v", v)
	}
	v := taxonomyViolation([]string{"geo", "rogue", "wild"}, taxonomy)
	if v == nil || v.Rule != "tag-taxonomy" || v.Detail != "tags not in geo.taxonomy: rogue, wild" {
		t.Errorf("violation = %+v", v)
	}
	if v := taxonomyViolation([]string{"anything"}, nil); v != nil {
		t.Errorf("undeclared taxonomy must skip: %+v", v)
	}
}
