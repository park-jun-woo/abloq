//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what filterOutlinks가 코퍼스 밖 대상과 자기 참조를 제외하는지 검증
package cluster

import (
	"reflect"
	"testing"
)

func TestFilterOutlinks(t *testing.T) {
	p := post{Section: "tech", Slug: "a", Outlinks: []string{"tech/a", "tech/b", "tech/gone", "tech/c"}}
	set := map[string]bool{"tech/a": true, "tech/b": true, "tech/c": true}
	got := filterOutlinks(p, set)
	want := []string{"tech/b", "tech/c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filterOutlinks = %v, want %v", got, want)
	}
}
