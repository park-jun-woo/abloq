//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what lineOf가 존재하는 키 경로는 그 라인을, 없는 경로는 1을 반환하는지 검증
package blogyaml

import "testing"

func TestLineOf(t *testing.T) {
	idx := lineIndex{"geo.freshness_days": 12}
	cases := []struct {
		name, path string
		want       int
	}{
		{"present key", "geo.freshness_days", 12},
		{"absent key falls back to 1", "geo.min_sources", 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := lineOf(idx, tc.path); got != tc.want {
				t.Errorf("lineOf(%q): want %d, got %d", tc.path, tc.want, got)
			}
		})
	}
}
