//ff:func feature=quest type=rule control=sequence
//ff:what 빌드 출력 축약 — 공백 정리 후 300자 절단, 비어 있으면 에러 문자열 (hugo-build Fact의 Actual)
package translation

import "strings"

// truncOut renders a failed build's combined output for Fact display.
func truncOut(out []byte, err error) string {
	s := strings.TrimSpace(string(out))
	if s == "" {
		s = err.Error()
	}
	if len(s) > 300 {
		s = s[:300] + "..."
	}
	return s
}
