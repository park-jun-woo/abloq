//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what sharedTags가 태그 교집합 크기를 계산하는지(빈 목록 포함) 검증
package cluster

import "testing"

func TestSharedTags(t *testing.T) {
	if got := sharedTags([]string{"geo", "abloq"}, []string{"abloq", "geo", "hugo"}); got != 2 {
		t.Errorf("sharedTags = %d, want 2", got)
	}
	if got := sharedTags([]string{"geo"}, []string{"hugo"}); got != 0 {
		t.Errorf("disjoint = %d, want 0", got)
	}
	if got := sharedTags(nil, []string{"geo"}); got != 0 {
		t.Errorf("nil = %d, want 0", got)
	}
}
