//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what edgeSet이 방향 간선 키("from->to")를 조립하는지 검증
package cluster

import "testing"

func TestEdgeSet(t *testing.T) {
	edges := edgeSet([]post{
		{Section: "tech", Slug: "a", Outlinks: []string{"tech/b"}},
		{Section: "tech", Slug: "b", Outlinks: []string{}},
	})
	if !edges["tech/a->tech/b"] || edges["tech/b->tech/a"] || len(edges) != 1 {
		t.Errorf("edgeSet = %v", edges)
	}
}
