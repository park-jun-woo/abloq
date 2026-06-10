//ff:func feature=visibility type=rule control=iteration dimension=1 topic=report
//ff:what WindowDates가 ym 월말 기준 양끝 포함 30일 윈도([월말-29, 월말])를 내는지 검증 — 2월·31일 월 경계 포함
package report

import "testing"

func TestWindowDates(t *testing.T) {
	cases := []struct{ ym, from, to string }{
		{"2026-04", "2026-04-01", "2026-04-30"}, // 30-day month = exactly the month
		{"2026-05", "2026-05-02", "2026-05-31"}, // 31-day month drops day 1
		{"2026-02", "2026-01-30", "2026-02-28"}, // short month reaches back into January
	}
	for _, tc := range cases {
		from, to, err := WindowDates(tc.ym)
		if err != nil {
			t.Fatalf("%s: %v", tc.ym, err)
		}
		if from != tc.from || to != tc.to {
			t.Errorf("%s: want [%s..%s], got [%s..%s]", tc.ym, tc.from, tc.to, from, to)
		}
	}
	if _, _, err := WindowDates(""); err == nil {
		t.Error("empty ym must error (resolve first)")
	}
}
