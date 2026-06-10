//ff:func feature=visibility type=scorer control=iteration dimension=1
//ff:what dateScore가 RFC3339·YYYY-MM-DD를 epoch 일수로 환산하고 미파싱 입력은 0을 반환하는지 검증
package priority

import "testing"

func TestDateScore(t *testing.T) {
	cases := []struct {
		date string
		want int64
	}{
		{"1970-01-02", 1},
		{"1970-01-02T12:00:00Z", 1},
		{"2026-06-03", 20607},
		{"not-a-date", 0},
		{"", 0},
	}
	for _, tc := range cases {
		if got := dateScore(tc.date); got != tc.want {
			t.Errorf("dateScore(%q): want %d, got %d", tc.date, tc.want, got)
		}
	}
}
