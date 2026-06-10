//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what 직전 닫힌 월(UTC) 산출 — 이번 달 1일 하루 전의 YYYY-MM, 기본 ym('')의 단일 정의
package report

import "time"

// LastClosedYM returns the previous month of now (UTC) as YYYY-MM — the
// single definition of the default ym. The matching SQL default lives in
// each window query's CASE branch; both must stay identical.
func LastClosedYM(now time.Time) string {
	u := now.UTC()
	first := time.Date(u.Year(), u.Month(), 1, 0, 0, 0, 0, time.UTC)
	return first.AddDate(0, 0, -1).Format("2006-01")
}
