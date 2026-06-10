//ff:func feature=archive type=client control=sequence
//ff:what processWayback이 target마다 영수증을 만들고 한 건의 실패가 나머지를 막지 않는지(부분 성공) 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessWayback(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		if strings.Contains(r.PostForm.Get("url"), "boom") {
			http.Error(w, "spn2 exploded", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write([]byte(`{"job_id":"ok"}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("WAYBACK_BASE_URL", srv.URL)

	items := processWayback([]Pending{
		{Kind: KindWayback, Target: "https://blog/a/"},
		{Kind: KindWayback, Target: "https://blog/boom/"},
		{Kind: KindWayback, Target: "https://blog/b/"},
	})
	if len(items) != 3 {
		t.Fatalf("len = %d, want 3", len(items))
	}
	if items[0].Status != StatusDone || items[2].Status != StatusDone {
		t.Errorf("healthy targets = %s/%s, want done/done", items[0].Status, items[2].Status)
	}
	if items[1].Status != StatusFailed {
		t.Errorf("boom target = %s, want failed (isolated)", items[1].Status)
	}
	if len(processWayback(nil)) != 0 {
		t.Error("empty group must produce no items")
	}
}
