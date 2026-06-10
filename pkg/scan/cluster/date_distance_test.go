//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what dateDistance가 발행일 거리의 절대값(일수)을 계산하는지 검증
package cluster

import "testing"

func TestDateDistance(t *testing.T) {
	cases := []struct {
		a, b string
		want int64
	}{
		{"2026-01-05", "2026-01-02", 3},
		{"2026-01-02", "2026-01-05", 3},
		{"2026-01-05", "2026-01-05", 0},
	}
	for _, tc := range cases {
		if got := dateDistance(tc.a, tc.b); got != tc.want {
			t.Errorf("dateDistance(%q, %q) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}
