//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what signRequest가 경로 없는 URL을 "/"로 정준화해 서명하는지 검증
package cflog

import (
	"net/http"
	"testing"
	"time"
)

func TestSignRequestEmptyPath(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://b.s3.us-east-1.amazonaws.com", nil)
	signRequest(req, "AKIDEXAMPLE", "secret", "", "us-east-1", "s3",
		time.Date(2015, 8, 30, 12, 36, 0, 0, time.UTC), emptyPayloadHash)
	if req.Header.Get("Authorization") == "" {
		t.Error("empty-path request not signed")
	}
}
