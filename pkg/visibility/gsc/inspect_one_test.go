//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what inspectOne이 1건 요약을 환원하고 비2xx·비JSON 응답이면 에러인지 검증
package gsc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInspectOne(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/fail":
			http.Error(w, `{"error":"quota"}`, http.StatusTooManyRequests)
		case "/badjson":
			_, _ = w.Write([]byte(`not-json`))
		default:
			_, _ = w.Write([]byte(`{"inspectionResult":{"indexStatusResult":{"verdict":"NEUTRAL","coverageState":"Discovered"}}}`))
		}
	}))
	defer srv.Close()

	ins, err := inspectOne(srv.URL+"/ok", "tok", "https://blog.test/", "https://blog.test/a/")
	if err != nil {
		t.Fatalf("inspectOne: %v", err)
	}
	if ins.URL != "https://blog.test/a/" || ins.Verdict != "NEUTRAL" || ins.CoverageState != "Discovered" {
		t.Errorf("inspection = %+v", ins)
	}

	if _, err := inspectOne(srv.URL+"/fail", "tok", "https://blog.test/", "u"); err == nil {
		t.Error("non-2xx must fail")
	}
	if _, err := inspectOne(srv.URL+"/badjson", "tok", "https://blog.test/", "u"); err == nil {
		t.Error("malformed JSON must fail")
	}
}
