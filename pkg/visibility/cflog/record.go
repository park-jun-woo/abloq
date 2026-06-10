//ff:type feature=visibility type=schema topic=crawl
//ff:what CF 로그 1행의 수집 필드 — UTC 시각, URL-디코드된 URI·UA, 상태 코드
package cflog

import "time"

// Record is one parsed CloudFront access-log line, reduced to the fields the
// crawl ingest consumes. URI and UA are already percent-decoded (CF encodes
// both; analyze-stats.py L59-61 applies the same unquote).
type Record struct {
	When   time.Time // UTC event time (date + time fields)
	URI    string    // cs-uri-stem, decoded
	Status string    // sc-status, verbatim
	UA     string    // cs(User-Agent), decoded
}
