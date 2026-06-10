//ff:func feature=visibility type=client control=iteration dimension=1 topic=gsc
//ff:what 증분 수집 1회 — 대상 일자를 오래된 순으로 Search Analytics 조회, 행 묶음과 일자 수 반환 (빈 일자 목록이면 no-op)
package gsc

// Collect runs one incremental GSC collection over the given closed days
// (from Dates). A mid-run failure aborts: the cursor derives from
// MAX(snap_date), so the next run resumes from the last fully stored day.
func Collect(base, token, site string, dates []string) (Result, error) {
	var res Result
	for _, date := range dates {
		rows, err := QueryDay(base, token, site, date)
		if err != nil {
			return Result{}, err
		}
		res.Rows = append(res.Rows, rows...)
		res.Days++
	}
	return res, nil
}
