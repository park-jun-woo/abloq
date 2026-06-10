//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what QueryDay가 사이트 경로 이스케이프·Bearer 헤더·1일 경계 바디로 조회하고 rows를 Snapshot으로 환원하는지, 비2xx면 에러인지 검증
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

func TestQueryDay(t *testing.T) {
	fail := false
	badJSON := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.EscapedPath(), "/webmasters/v3/sites/https:%2F%2Fblog.test%2F/searchAnalytics/query") {
			t.Errorf("path = %q", r.URL.EscapedPath())
		}
		if r.Header.Get("Authorization") != "Bearer tok-1" {
			t.Errorf("authorization = %q", r.Header.Get("Authorization"))
		}
		body, _ := io.ReadAll(r.Body)
		var req map[string]any
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("request body: %v", err)
		}
		if req["startDate"] != "2026-06-08" || req["endDate"] != "2026-06-08" {
			t.Errorf("date bounds = %v / %v", req["startDate"], req["endDate"])
		}
		if fail {
			http.Error(w, `{"error":"denied"}`, http.StatusForbidden)
			return
		}
		if badJSON {
			_, _ = w.Write([]byte(`not-json`))
			return
		}
		_, _ = w.Write([]byte(`{"rows":[
			{"keys":["https://blog.test/tech/post-a/"],"clicks":3,"impressions":120,"ctr":0.025,"position":4.2},
			{"keys":[],"clicks":1,"impressions":1,"position":1},
			{"keys":["https://blog.test/tech/post-b/"],"clicks":0,"impressions":40,"ctr":0,"position":9.8}
		]}`))
	}))
	defer srv.Close()

	rows, err := QueryDay(srv.URL, "tok-1", "https://blog.test/", "2026-06-08")
	if err != nil {
		t.Fatalf("QueryDay: %v", err)
	}
	want := []Snapshot{
		{SnapDate: "2026-06-08", Page: "https://blog.test/tech/post-a/", Impressions: 120, Clicks: 3, AvgPosition: 4.2},
		{SnapDate: "2026-06-08", Page: "https://blog.test/tech/post-b/", Impressions: 40, Clicks: 0, AvgPosition: 9.8},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %+v, want %+v", rows, want)
	}

	fail = true
	if _, err := QueryDay(srv.URL, "tok-1", "https://blog.test/", "2026-06-08"); err == nil {
		t.Error("non-2xx must fail")
	}
	fail, badJSON = false, true
	if _, err := QueryDay(srv.URL, "tok-1", "https://blog.test/", "2026-06-08"); err == nil {
		t.Error("malformed JSON must fail")
	}
}
