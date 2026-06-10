//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what directions가 기존재 간선 방향을 제외하고 out/in을 산출하는지(완전 연결 = 빈 목록) 검증
package cluster

import (
	"reflect"
	"testing"
)

func TestDirections(t *testing.T) {
	a := post{Section: "tech", Slug: "a"}
	b := post{Section: "tech", Slug: "b"}
	cases := []struct {
		name  string
		edges map[string]bool
		want  []string
	}{
		{"no edges", map[string]bool{}, []string{"out", "in"}},
		{"out exists", map[string]bool{"tech/a->tech/b": true}, []string{"in"}},
		{"in exists", map[string]bool{"tech/b->tech/a": true}, []string{"out"}},
		{"fully connected", map[string]bool{"tech/a->tech/b": true, "tech/b->tech/a": true}, []string{}},
	}
	for _, tc := range cases {
		if got := directions(a, b, tc.edges); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: directions = %v, want %v", tc.name, got, tc.want)
		}
	}
}
