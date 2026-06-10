//ff:func feature=visibility type=client control=iteration dimension=1 topic=gsc
//ff:what 수집 대상 일자 목록 계산 — 커서 다음 날부터 오늘(UTC)−지연마진까지의 닫힌 일자만, 첫 수집(빈 커서)은 lookback일
//ff:why GSC 데이터는 2~3일 지연 확정된다 — 마진 안의 열린 일자를 적재하면 미확정 수치가 굳는다. 커서는 MAX(snap_date) 파생: 같은 일자 재수집 중복 0을 커서가 보증한다 (Phase013)
package gsc

import "time"

// Dates returns the closed UTC days one GSC collection must cover, oldest
// first: from the day after the cursor (MAX(snap_date), "" on first run —
// then lookback days) up to today minus the delay margin, inclusive. An
// up-to-date cursor yields nil.
func Dates(cursor string, today time.Time, marginDays, lookbackDays int) []string {
	end := today.UTC().AddDate(0, 0, -marginDays)
	start := end.AddDate(0, 0, -(lookbackDays - 1))
	if cursor != "" {
		cur, err := time.Parse("2006-01-02", cursor)
		if err != nil {
			return nil
		}
		start = cur.AddDate(0, 0, 1)
	}
	var dates []string
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}
	return dates
}
