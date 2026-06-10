//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what S3Source.List가 비-2xx 응답과 깨진 XML을 에러로 전파하는지 검증 (start-after 경로 포함)
package cflog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestS3ListErrors(t *testing.T) {
	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("prefix") == "logs/deny" {
			http.Error(w, "denied", http.StatusForbidden)
			return
		}
		fmt.Fprint(w, "this is not xml <<<")
	}))
	defer bad.Close()
	src := S3Source{
		Bucket: "b", Prefix: "logs/", Region: "us-east-1",
		AccessKey: "k", SecretKey: "s", Endpoint: bad.URL, Client: bad.Client(),
	}
	if _, err := src.List("deny", ""); err == nil {
		t.Error("403 list accepted")
	}
	if _, err := src.List("", "after-key"); err == nil {
		t.Error("broken XML accepted (start-after path)")
	}
}
