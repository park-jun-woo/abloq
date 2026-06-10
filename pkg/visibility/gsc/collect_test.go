//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what Collect가 일자별 조회를 합산(행 누적·Days 카운트)하고 빈 일자 목록이면 no-op, 중도 실패면 에러로 중단하는지 검증
package gsc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCollect(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls > 2 {
			http.Error(w, `{"error":"boom"}`, http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte(`{"rows":[{"keys":["https://blog.test/p/"],"clicks":1,"impressions":10,"position":2.5}]}`))
	}))
	defer srv.Close()

	res, err := Collect(srv.URL, "tok", "https://blog.test/", []string{"2026-06-07", "2026-06-08"})
	if err != nil {
		t.Fatalf("Collect: %v", err)
	}
	if res.Days != 2 || len(res.Rows) != 2 {
		t.Errorf("days=%d rows=%d, want 2/2", res.Days, len(res.Rows))
	}
	if res.Rows[0].SnapDate != "2026-06-07" || res.Rows[1].SnapDate != "2026-06-08" {
		t.Errorf("snap dates = %q, %q", res.Rows[0].SnapDate, res.Rows[1].SnapDate)
	}

	if res, err := Collect(srv.URL, "tok", "https://blog.test/", nil); err != nil || res.Days != 0 || res.Rows != nil {
		t.Errorf("empty dates = %+v, %v, want zero no-op", res, err)
	}

	if _, err := Collect(srv.URL, "tok", "https://blog.test/", []string{"2026-06-09"}); err == nil {
		t.Error("mid-run failure must abort with an error")
	}
}
