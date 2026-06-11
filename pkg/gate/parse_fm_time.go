//ff:func feature=gate type=parser control=selection
//ff:what front matter 날짜 값(time.Time 또는 RFC3339/날짜 문자열)을 time.Time으로 해석 (공개 API)
//ff:why Phase017 번역 퀘스트의 fm-mirror(⑦)가 원문↔번역 date·lastmod를 같은 해석기로 비교해야 해서 export — 재구현(복제) 대신 단일 출처 유지
package gate

import "time"

// ParseFMTime interprets a decoded front matter date value. Exported for the
// translation quest's fm-mirror comparison (Phase017).
func ParseFMTime(v any) (time.Time, bool) {
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
