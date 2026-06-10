//ff:type feature=visibility type=schema topic=crawl
//ff:what 수집 1회의 산출 — 히트 행, 전진된 커서, 미지 봇 행, 무필터 원시 봇 카운터, 처리 파일·히트 합계
package cflog

// Result is one crawl ingest outcome. Hits/Cursors/Unknown feed the three
// batch upserts; Raw is the per-bot counter over every parsed line with no
// status or mapping filter (the analyze-stats.py comparison point); Files is
// the number of log objects processed and Total the sum of Hits+MDHits over
// all rows.
type Result struct {
	Hits    []HitRow
	Cursors []Cursor
	Unknown []UnknownRow
	Raw     map[string]int64
	Files   int64
	Total   int64
}
