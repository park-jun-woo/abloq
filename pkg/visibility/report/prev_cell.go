//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what 전월 표 셀 — 전월 데이터가 없으면 "n/a"(첫 달), 있으면 값
package report

import "strconv"

// prevCell renders one previous-month table cell: "n/a" when the previous
// window has no data at all (the first month), the value otherwise.
func prevCell(r Report, v int64) string {
	if !r.PrevHasData {
		return "n/a"
	}
	return strconv.FormatInt(v, 10)
}
