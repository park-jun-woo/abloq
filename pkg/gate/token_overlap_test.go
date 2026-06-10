//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what tokenOverlap 케이스 — 전부 포함 1.0, 부분 포함 비율, 무관 0, 빈 want는 1
package gate

import "testing"

func TestTokenOverlap(t *testing.T) {
	cases := []struct {
		name      string
		want, got []string
		wantRatio float64
	}{
		{"full overlap", []string{"a", "b"}, []string{"b", "a", "c"}, 1},
		{"half overlap", []string{"a", "b"}, []string{"a", "x"}, 0.5},
		{"no overlap", []string{"a"}, []string{"x"}, 0},
		{"empty want", nil, []string{"x"}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tokenOverlap(tc.want, tc.got); got != tc.wantRatio {
				t.Errorf("tokenOverlap = %v, want %v", got, tc.wantRatio)
			}
		})
	}
}
