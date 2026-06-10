//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what PrevYM이 연 경계를 포함해 한 달 전 ym을 내고 미파싱 입력은 빈 문자열인지 검증
package report

import "testing"

func TestPrevYM(t *testing.T) {
	if got := PrevYM("2026-04"); got != "2026-03" {
		t.Errorf("want 2026-03, got %q", got)
	}
	if got := PrevYM("2026-01"); got != "2025-12" {
		t.Errorf("year boundary: want 2025-12, got %q", got)
	}
	if got := PrevYM("nope"); got != "" {
		t.Errorf("unparseable ym must yield empty: %q", got)
	}
}
