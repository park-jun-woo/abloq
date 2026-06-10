//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what TokenDiff가 추가+삭제 토큰 수(multiset 대칭차)를 계산하는지 검증
package gate

import "testing"

func TestTokenDiff(t *testing.T) {
	cases := []struct {
		name string
		a, b []string
		want int
	}{
		{"identical", []string{"a", "b"}, []string{"a", "b"}, 0},
		{"reordered", []string{"a", "b"}, []string{"b", "a"}, 0},
		{"added two", []string{"a"}, []string{"a", "x", "y"}, 2},
		{"removed one", []string{"a", "b"}, []string{"a"}, 1},
		{"replaced", []string{"a"}, []string{"b"}, 2},
		{"multiplicity", []string{"a", "a"}, []string{"a"}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TokenDiff(tc.a, tc.b); got != tc.want {
				t.Errorf("TokenDiff = %d, want %d", got, tc.want)
			}
		})
	}
}
