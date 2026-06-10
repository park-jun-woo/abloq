//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what inlinkCounts가 확정 아웃링크의 역방향 합으로 인링크 수를 집계하는지 검증
package cluster

import "testing"

func TestInlinkCounts(t *testing.T) {
	counts := inlinkCounts([]post{
		{Section: "tech", Slug: "a", Outlinks: []string{"tech/b", "tech/c"}},
		{Section: "tech", Slug: "b", Outlinks: []string{"tech/c"}},
		{Section: "tech", Slug: "c", Outlinks: []string{}},
	})
	if counts["tech/c"] != 2 || counts["tech/b"] != 1 || counts["tech/a"] != 0 {
		t.Errorf("inlinkCounts = %v", counts)
	}
}
