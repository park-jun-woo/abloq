//ff:func feature=visibility type=scorer control=iteration dimension=1
//ff:what ColdStart.Score가 hits>0이면 hits를, 아니면 date 점수를 반환하고 date 역순(최신 우선) 순서를 보존하는지 검증
package priority

import "testing"

func TestColdStartScore(t *testing.T) {
	cases := []struct {
		name string
		s    Signals
		want int64
	}{
		{"hits win", Signals{Date: "2026-06-01", Hits: 7}, 7},
		{"date fallback", Signals{Date: "1970-01-03", Hits: 0}, 2},
		{"empty date", Signals{Date: "", Hits: 0}, 0},
	}
	for _, tc := range cases {
		if got := (ColdStart{}).Score(tc.s); got != tc.want {
			t.Errorf("%s: want %d, got %d", tc.name, tc.want, got)
		}
	}
	older := (ColdStart{}).Score(Signals{Date: "2026-06-01"})
	newer := (ColdStart{}).Score(Signals{Date: "2026-06-03"})
	if newer <= older {
		t.Errorf("newer date must score higher: %d <= %d", newer, older)
	}
}
