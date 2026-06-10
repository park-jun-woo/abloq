//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what postSet이 노드 키 집합을 조립하는지 검증
package cluster

import "testing"

func TestPostSet(t *testing.T) {
	set := postSet([]post{{Section: "tech", Slug: "a"}, {Section: "opinion", Slug: "a"}})
	if len(set) != 2 || !set["tech/a"] || !set["opinion/a"] {
		t.Errorf("postSet = %v", set)
	}
}
