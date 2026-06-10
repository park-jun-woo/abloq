//ff:func feature=scan type=parser control=iteration dimension=1
//ff:what parseDate가 RFC3339·YYYY-MM-DD를 받고 그 외는 ok=false인지 검증
package freshness

import "testing"

func TestParseDate(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"2026-06-05", true},
		{"2026-06-05T09:00:00+09:00", true},
		{"2026/06/05", false},
		{"", false},
	}
	for _, tc := range cases {
		if _, ok := parseDate(tc.in); ok != tc.ok {
			t.Errorf("parseDate(%q): want ok=%v, got %v", tc.in, tc.ok, ok)
		}
	}
}
