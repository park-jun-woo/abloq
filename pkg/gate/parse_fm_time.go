//ff:func feature=gate type=parser control=selection
//ff:what front matter 날짜 값(time.Time 또는 RFC3339/날짜 문자열)을 time.Time으로 해석
package gate

import "time"

// parseFMTime interprets a decoded front matter date value.
func parseFMTime(v any) (time.Time, bool) {
	switch d := v.(type) {
	case time.Time:
		return d, true
	case string:
		if ts, err := time.Parse(time.RFC3339, d); err == nil {
			return ts, true
		}
		if ts, err := time.Parse("2006-01-02", d); err == nil {
			return ts, true
		}
	}
	return time.Time{}, false
}
