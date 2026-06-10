//ff:func feature=scan type=parser control=iteration dimension=1 topic=cluster
//ff:what epochDay가 RFC3339·YYYY-MM-DD를 epoch 일수로 환산하고 미파싱 입력은 0인지 검증
package cluster

import "testing"

func TestEpochDay(t *testing.T) {
	cases := []struct {
		date string
		want int64
	}{
		{"1970-01-02", 1},
		{"1970-01-03T09:00:00+09:00", 2},
		{"not-a-date", 0},
		{"", 0},
	}
	for _, tc := range cases {
		if got := epochDay(tc.date); got != tc.want {
			t.Errorf("epochDay(%q) = %d, want %d", tc.date, got, tc.want)
		}
	}
}
