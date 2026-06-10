//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 증분 수집 1회 — 커서 시간대 초과·마지막 닫힌 시간대 이하의 키만 통째로 수집하고 커서를 시간 경계로 전진
//ff:why 커서는 시간 경계 단위: 같은 로그 재수집 중복 0을 커서가 보증한다. 커서 리셋 재집계는 반드시 대상 구간 crawl_hits 삭제와 한 동작(리셋만 하면 이중 누적 — 운영 절차, 수동 SQL). afterKey 리스팅은 쓰지 않는다(start-after 금지) (Phase012)
package cflog

import "time"

// Collect runs one incremental crawl ingest: list every key, keep only the
// hours in (cursor, lastClosed(now, margin)], aggregate them wholesale and
// advance the cursor to the closed-hour boundary. Hours inside the safety
// margin stay untouched so late-delivered files are not lost.
func Collect(src Source, urls map[string]Article, cursors []Cursor, now time.Time, margin time.Duration) (Result, error) {
	keys, err := src.List("", "")
	if err != nil {
		return Result{}, err
	}
	closed := lastClosedHour(now, margin)
	selected := selectKeys(keys, cursorHourFor(cursors, CursorSource), closed)
	agg, err := IngestKeys(src, urls, selected)
	if err != nil {
		return Result{}, err
	}
	rows := agg.HitRows()
	return Result{
		Hits:    rows,
		Cursors: []Cursor{{Source: CursorSource, CursorHour: closed}},
		Unknown: agg.UnknownRows(),
		Raw:     agg.Raw,
		Files:   int64(len(selected)),
		Total:   TotalHits(rows),
	}, nil
}
