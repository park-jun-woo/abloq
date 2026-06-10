//ff:func feature=gate type=rule control=iteration dimension=1 topic=lossless
//ff:what MultisetSubset이 중복도까지 비교(부분집합/재배열 허용, 삭제·중복 감소 검출)하는지 검증
package gate

import "testing"

func TestMultisetSubset(t *testing.T) {
	cases := []struct {
		name        string
		want, have  []string
		wantMissing string
		wantOK      bool
	}{
		{"equal", []string{"a", "b"}, []string{"a", "b"}, "", true},
		{"reordered", []string{"a", "b"}, []string{"b", "a"}, "", true},
		{"superset ok", []string{"a"}, []string{"a", "x"}, "", true},
		{"deleted", []string{"a", "b"}, []string{"a"}, "b", false},
		{"multiplicity drop", []string{"a", "a"}, []string{"a"}, "a", false},
		{"empty want", nil, nil, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			miss, ok := MultisetSubset(tc.want, tc.have)
			if ok != tc.wantOK || miss != tc.wantMissing {
				t.Errorf("got (%q, %v), want (%q, %v)", miss, ok, tc.wantMissing, tc.wantOK)
			}
		})
	}
}
