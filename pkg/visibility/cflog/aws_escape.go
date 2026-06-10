//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what RFC3986 퍼센트 인코딩 — 비예약 문자(A-Za-z0-9-._~)만 그대로, keepSlash면 경로 구분자 '/'도 유지
package cflog

import "fmt"

// awsEscape percent-encodes s the way SigV4 canonicalization requires:
// unreserved characters stay bare, everything else becomes uppercase %XX.
// keepSlash leaves '/' intact for path encoding.
func awsEscape(s string, keepSlash bool) string {
	var out []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
			c == '-' || c == '.' || c == '_' || c == '~' || (keepSlash && c == '/') {
			out = append(out, c)
			continue
		}
		out = append(out, []byte(fmt.Sprintf("%%%02X", c))...)
	}
	return string(out)
}
