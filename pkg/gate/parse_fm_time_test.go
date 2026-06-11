//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what parseFMTime이 time.Time/RFC3339 문자열/날짜 문자열을 해석하고 그 외를 거부하는지 검증
package gate

import (
	"testing"
	"time"
)

func TestParseFMTime(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		name   string
		in     any
		wantOK bool
	}{
		{"time.Time", now, true},
		{"rfc3339", "2026-01-01T09:00:00+09:00", true},
		{"date only", "2026-01-01", true},
		{"garbage", "not-a-date", false},
		{"nil", nil, false},
		{"number", 42, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, ok := ParseFMTime(tc.in); ok != tc.wantOK {
				t.Errorf("ParseFMTime(%v) ok = %v, want %v", tc.in, ok, tc.wantOK)
			}
		})
	}
}
