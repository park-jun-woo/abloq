//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what cursorHourFor가 소스 이름의 커서를 찾고 없으면 빈 문자열인지 검증
package cflog

import "testing"

func TestCursorHourFor(t *testing.T) {
	cursors := []Cursor{{Source: "other", CursorHour: "2026-01-01-00"}, {Source: CursorSource, CursorHour: "2026-06-01-23"}}
	if got := cursorHourFor(cursors, CursorSource); got != "2026-06-01-23" {
		t.Errorf("cursorHourFor = %q", got)
	}
	if got := cursorHourFor(nil, CursorSource); got != "" {
		t.Errorf("cursorHourFor(nil) = %q, want empty", got)
	}
}
