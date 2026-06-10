//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what checkHost가 한 호스트의 URL들을 순차 점검해 url→판정 맵을 채우는지 검증 (지연 0으로 고정)
package evidence

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/dead" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	c := &Checker{Client: srv.Client(), UserAgent: "abloqd-linkcheck"}
	res := c.checkHost([]string{srv.URL + "/alive", srv.URL + "/dead"})
	if res[srv.URL+"/alive"] != "ok" || res[srv.URL+"/dead"] != "hard" {
		t.Errorf("checkHost = %v", res)
	}
}
