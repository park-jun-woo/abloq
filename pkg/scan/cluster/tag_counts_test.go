//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what tagCounts가 태그별 보유 글 수를 코퍼스 전체에서 집계하는지 검증
package cluster

import "testing"

func TestTagCounts(t *testing.T) {
	counts := tagCounts([]post{
		{Tags: []string{"geo", "abloq"}},
		{Tags: []string{"geo"}},
		{Tags: []string{}},
	})
	if counts["geo"] != 2 || counts["abloq"] != 1 || len(counts) != 2 {
		t.Errorf("tagCounts = %v", counts)
	}
}
