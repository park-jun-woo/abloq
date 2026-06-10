//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what s3EscapePath가 '/' 구분자를 남기고 키를 경로 인코딩하는지 검증
package cflog

import "testing"

func TestS3EscapePath(t *testing.T) {
	if got := s3EscapePath("logs/E.2026-06-01-12.a b.gz"); got != "logs/E.2026-06-01-12.a%20b.gz" {
		t.Errorf("s3EscapePath = %q", got)
	}
}
