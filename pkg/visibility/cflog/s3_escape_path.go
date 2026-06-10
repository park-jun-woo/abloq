//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what S3 객체 키를 URL 경로로 인코딩 — 비예약 문자와 '/'만 유지
package cflog

// s3EscapePath percent-encodes an object key for the request path, keeping
// the '/' separators.
func s3EscapePath(key string) string {
	return awsEscape(key, true)
}
