//ff:func feature=visibility type=rule control=sequence topic=report
//ff:what ym 인자 해석 — ''이면 직전 닫힌 월(UTC), 명시 값은 YYYY-MM 형식 검증 (명시 경로에 now 미개입)
package report

import (
	"fmt"
	"time"
)

// ResolveYM resolves the request's ym argument: the empty string means the
// last closed month (the only place now enters), an explicit value must be
// YYYY-MM and passes through untouched — the golden tests and Hurl pin
// explicit months precisely because no clock is consulted on that path.
func ResolveYM(ym string, now time.Time) (string, error) {
	if ym == "" {
		return LastClosedYM(now), nil
	}
	if _, err := time.Parse("2006-01", ym); err != nil {
		return "", fmt.Errorf("invalid ym %q: want YYYY-MM", ym)
	}
	return ym, nil
}
