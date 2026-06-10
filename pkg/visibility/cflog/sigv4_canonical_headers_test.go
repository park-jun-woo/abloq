//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what canonicalHeaders가 host 포함 전 헤더를 소문자 정렬·트림해 정준 블록과 서명 목록으로 푸는지 검증
package cflog

import (
	"net/http"
	"testing"
)

func TestCanonicalHeaders(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://iam.amazonaws.com/", nil)
	req.Header.Set("X-Amz-Date", "20150830T123600Z")
	req.Header.Set("Content-Type", " application/x-www-form-urlencoded; charset=utf-8 ")
	block, signed := canonicalHeaders(req)
	wantBlock := "content-type:application/x-www-form-urlencoded; charset=utf-8\n" +
		"host:iam.amazonaws.com\n" +
		"x-amz-date:20150830T123600Z\n"
	if block != wantBlock {
		t.Errorf("block = %q, want %q", block, wantBlock)
	}
	if signed != "content-type;host;x-amz-date" {
		t.Errorf("signed = %q", signed)
	}
}
