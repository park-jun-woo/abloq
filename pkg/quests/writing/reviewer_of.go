//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what REVIEW 기록에서 "reviewer:" 라인의 검토 컨텍스트 식별자 추출 — 없거나 값이 비면 빈 문자열
package writing

import "strings"

// reviewerOf extracts the reviewer context identifier from a review record:
// the value of the first non-empty `reviewer:` line.
func reviewerOf(review string) string {
	for _, line := range strings.Split(review, "\n") {
		rest, ok := strings.CutPrefix(strings.TrimSpace(line), "reviewer:")
		if !ok {
			continue
		}
		if v := strings.TrimSpace(rest); v != "" {
			return v
		}
	}
	return ""
}
