//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what sortedKeys가 맵 키를 오름차순으로 반환하는지(빈 맵 포함) 검증
package blogyaml

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	cases := []struct {
		name string
		m    map[string]string
		want []string
	}{
		{"empty map", map[string]string{}, []string{}},
		{"unsorted keys", map[string]string{"zeta": "1", "alpha": "2", "mid": "3"}, []string{"alpha", "mid", "zeta"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := sortedKeys(tc.m); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
