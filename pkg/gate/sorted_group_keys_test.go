//ff:func feature=gate type=rule control=sequence
//ff:what sortedGroupKeys가 그룹 키를 오름차순으로 반환하는지 검증
package gate

import (
	"reflect"
	"testing"
)

func TestSortedGroupKeys(t *testing.T) {
	groups := map[string][]*Article{"tech/b": nil, "tech/a": nil, "opinion/c": nil}
	got := sortedGroupKeys(groups)
	want := []string{"opinion/c", "tech/a", "tech/b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedGroupKeys = %v, want %v", got, want)
	}
}
