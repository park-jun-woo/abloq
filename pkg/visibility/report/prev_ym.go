//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what 전월 ym 산출 — YYYY-MM의 한 달 전 (연 경계 안전: 1일 기준 가감)
package report

import "time"

// PrevYM returns the month before ym. The arithmetic anchors on the first
// of the month, so year boundaries and month-length differences are safe.
// ym must already be validated (ResolveYM); an unparseable value returns "".
func PrevYM(ym string) string {
	t, err := time.Parse("2006-01", ym)
	if err != nil {
		return ""
	}
	return t.AddDate(0, -1, 0).Format("2006-01")
}
