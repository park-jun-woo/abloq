//ff:type feature=visibility type=schema topic=crawl
//ff:what 수집 커서 1건 — 소스 이름 + 마지막으로 통째로 수집한 닫힌 시간대(YYYY-MM-DD-HH, UTC), ingest_cursors 1행 대응
//ff:why 커서는 시간 경계 단위다: CF 키의 시간 뒤 접미사가 랜덤이라 키 기반 start-after 커서는 지연 배달 파일을 영구 누락시킨다(과소집계) — JSON 키는 ingest_cursors 컬럼명과 일치해야 백엔드 jsonb 공급/업서트가 그대로 통한다 (Phase012)
package cflog

// CursorSource is the ingest_cursors source name of the CloudFront crawl
// ingest.
const CursorSource = "cf_logs"

// Cursor is one ingest cursor row: every hour up to and including CursorHour
// (UTC, "YYYY-MM-DD-HH") has been ingested wholesale. Empty CursorHour means
// nothing was ingested yet. The string format sorts chronologically.
type Cursor struct {
	Source     string `json:"source"`
	CursorHour string `json:"cursor_hour"`
}
