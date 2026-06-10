//ff:func feature=visibility type=rule control=sequence topic=report
//ff:what ym → 30d 윈도 경계 — ym 월말 역산 30일, 양끝 포함 [월말-29, 월말] (SQL CASE 윈도와 동일 정의, now 미사용)
package report

import (
	"fmt"
	"time"
)

// WindowDates returns the report window for ym: the 30 days ending on the
// month's last day, both endpoints included ([month_end-29 .. month_end]).
// The SQL window queries implement the identical bounds — the golden tests
// and the CLI's log filter both depend on the two definitions agreeing.
func WindowDates(ym string) (string, string, error) {
	t, err := time.Parse("2006-01", ym)
	if err != nil {
		return "", "", fmt.Errorf("invalid ym %q: want YYYY-MM", ym)
	}
	end := t.AddDate(0, 1, -1)
	return end.AddDate(0, 0, -29).Format("2006-01-02"), end.Format("2006-01-02"), nil
}
