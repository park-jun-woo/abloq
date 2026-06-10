//ff:func feature=archive type=client control=sequence
//ff:what saveWayback이 url 폼을 제출하고 2xx면 done·5xx면 failed·연결 불가면 failed(status_code 0)인지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSaveWayback(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/save" {
			t.Errorf("path = %s, want /save", r.URL.Path)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		if strings.Contains(r.PostForm.Get("url"), "boom") {
			http.Error(w, "spn2 exploded", http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"url":"` + r.PostForm.Get("url") + `","job_id":"j-1"}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("WAYBACK_BASE_URL", srv.URL)

	ok := saveWayback(Pending{DeployID: "d", Kind: KindWayback, Target: "https://blog/p/"})
	if ok.Status != StatusDone || !strings.Contains(string(ok.Response), "j-1") {
		t.Errorf("2xx save = %+v, want done with job_id evidence", ok)
	}
	bad := saveWayback(Pending{DeployID: "d", Kind: KindWayback, Target: "https://blog/boom/"})
	if bad.Status != StatusFailed || !strings.Contains(string(bad.Response), "502") {
		t.Errorf("5xx save = %+v, want failed with 502 evidence", bad)
	}

	t.Setenv("WAYBACK_BASE_URL", "http://127.0.0.1:1")
	dead := saveWayback(Pending{DeployID: "d", Kind: KindWayback, Target: "https://blog/p/"})
	if dead.Status != StatusFailed || !strings.Contains(string(dead.Response), `"status_code":0`) {
		t.Errorf("transport error = %+v, want failed with status_code 0", dead)
	}
}
