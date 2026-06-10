//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 커서 목록에서 소스 이름의 커서 시간대 조회 — 없으면 빈 문자열(아직 아무것도 수집 안 함)
package cflog

// cursorHourFor returns the cursor hour recorded for one source, or "" when
// the source has no cursor row yet.
func cursorHourFor(cursors []Cursor, source string) string {
	for _, c := range cursors {
		if c.Source == source {
			return c.CursorHour
		}
	}
	return ""
}
