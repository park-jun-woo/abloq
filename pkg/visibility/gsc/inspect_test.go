//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what Inspect가 URL마다 inspectionUrl·siteUrl 바디로 1건씩 조회하고 verdict·coverageState 요약을 환원하는지 검증
package gsc

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestInspect(t *testing.T) {
	var seen []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/v1/urlInspection/index:inspect") {
			t.Errorf("path = %q", r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		var req map[string]string
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("request body: %v", err)
		}
		if req["siteUrl"] != "https://blog.test/" {
			t.Errorf("siteUrl = %q", req["siteUrl"])
		}
		if strings.Contains(req["inspectionUrl"], "boom") {
			http.Error(w, `{"error":"boom"}`, http.StatusInternalServerError)
			return
		}
		seen = append(seen, req["inspectionUrl"])
		_, _ = w.Write([]byte(`{"inspectionResult":{"indexStatusResult":{"verdict":"PASS","coverageState":"Submitted and indexed"}}}`))
	}))
	defer srv.Close()

	urls := []string{"https://blog.test/tech/a/", "https://blog.test/tech/b/"}
	got, err := Inspect(srv.URL, "tok", "https://blog.test/", urls)
	if err != nil {
		t.Fatalf("Inspect: %v", err)
	}
	want := []Inspection{
		{URL: urls[0], Verdict: "PASS", CoverageState: "Submitted and indexed"},
		{URL: urls[1], Verdict: "PASS", CoverageState: "Submitted and indexed"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("inspections = %+v, want %+v", got, want)
	}
	if !reflect.DeepEqual(seen, urls) {
		t.Errorf("inspected URLs = %v, want %v", seen, urls)
	}

	if out, err := Inspect(srv.URL, "tok", "https://blog.test/", nil); err != nil || out != nil {
		t.Errorf("no URLs = %v, %v, want nil no-op", out, err)
	}
	if _, err := Inspect(srv.URL, "tok", "https://blog.test/", []string{"https://blog.test/boom/"}); err == nil {
		t.Error("mid-run failure must abort with an error")
	}
}
