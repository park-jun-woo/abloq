//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what CheckAll이 호스트 2개의 URL들을 병렬 점검해 전체 url→판정 맵을 돌려주는지 검증 (Concurrency 0도 동작)
package evidence

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckAll(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/dead" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	srvA := httptest.NewServer(handler)
	defer srvA.Close()
	srvB := httptest.NewServer(handler)
	defer srvB.Close()
	c := &Checker{Client: http.DefaultClient, UserAgent: "abloqd-linkcheck"} // Concurrency 0 → min 1
	res := c.CheckAll([]string{srvA.URL + "/alive", srvA.URL + "/dead", srvB.URL + "/alive"})
	if len(res) != 3 {
		t.Fatalf("CheckAll = %v, want 3 verdicts", res)
	}
	if res[srvA.URL+"/alive"] != "ok" || res[srvA.URL+"/dead"] != "hard" || res[srvB.URL+"/alive"] != "ok" {
		t.Errorf("CheckAll = %v", res)
	}
}
