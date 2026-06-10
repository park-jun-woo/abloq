//ff:func feature=visibility type=rule control=sequence topic=report
//ff:what ResolveYM이 ''→직전 닫힌 월(UTC), 명시 YYYY-MM은 무변형 통과, 형식 위반은 에러인지 검증
package report

import (
	"testing"
	"time"
)

func TestResolveYM(t *testing.T) {
	now := time.Date(2026, 6, 11, 3, 0, 0, 0, time.UTC)
	got, err := ResolveYM("", now)
	if err != nil || got != "2026-05" {
		t.Errorf("empty ym: want 2026-05, got %q (%v)", got, err)
	}
	// January boundary: the previous closed month crosses the year.
	got, err = ResolveYM("", time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC))
	if err != nil || got != "2025-12" {
		t.Errorf("january: want 2025-12, got %q (%v)", got, err)
	}
	got, err = ResolveYM("2026-04", now)
	if err != nil || got != "2026-04" {
		t.Errorf("explicit ym must pass through: %q (%v)", got, err)
	}
	if _, err := ResolveYM("2026-4", now); err == nil {
		t.Error("malformed ym must error")
	}
	if _, err := ResolveYM("not-a-ym", now); err == nil {
		t.Error("garbage ym must error")
	}
}
