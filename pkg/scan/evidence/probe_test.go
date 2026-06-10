//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what probe가 UA 헤더를 달고 최종 상태 코드를 회수하는지, 요청 생성 불가 URL이 에러인지 검증
package evidence

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProbe(t *testing.T) {
	var gotUA, gotMethod string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA, gotMethod = r.UserAgent(), r.Method
		w.WriteHeader(http.StatusTeapot)
	}))
	defer srv.Close()
	c := &Checker{Client: srv.Client(), UserAgent: "abloqd-linkcheck"}
	code, err := c.probe(http.MethodHead, srv.URL+"/x")
	if err != nil || code != http.StatusTeapot {
		t.Fatalf("probe = (%d, %v), want (418, nil)", code, err)
	}
	if gotUA != "abloqd-linkcheck" || gotMethod != "HEAD" {
		t.Errorf("request was (%q, %q), want declared UA and HEAD", gotUA, gotMethod)
	}
	if _, err := c.probe("BAD METHOD", srv.URL); err == nil {
		t.Error("invalid request must error")
	}
}
