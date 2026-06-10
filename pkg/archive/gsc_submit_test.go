//ff:func feature=archive type=client control=sequence
//ff:what gscSubmit이 2xx done·429 failed(quota_exceeded 마커)·5xx failed·연결 불가 failed를 매기는지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGscSubmit(t *testing.T) {
	mode := http.StatusOK
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("Authorization = %q", r.Header.Get("Authorization"))
		}
		w.WriteHeader(mode)
		if _, err := w.Write([]byte(`{"urlNotificationMetadata":{}}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	endpoint := srv.URL + "/v3/urlNotifications:publish"
	p := Pending{DeployID: "d", Kind: KindGSCIndex, Target: "https://blog/p/"}

	ok := gscSubmit(p, endpoint, "tok")
	if ok.Status != StatusDone {
		t.Errorf("2xx = %s, want done", ok.Status)
	}

	mode = http.StatusTooManyRequests
	quota := gscSubmit(p, endpoint, "tok")
	if quota.Status != StatusFailed || !strings.Contains(string(quota.Response), `"quota_exceeded":true`) {
		t.Errorf("429 = %+v, want failed with quota_exceeded marker", quota)
	}

	mode = http.StatusBadRequest
	bad := gscSubmit(p, endpoint, "tok")
	if bad.Status != StatusFailed || strings.Contains(string(bad.Response), "quota_exceeded") {
		t.Errorf("4xx = %+v, want plain failed", bad)
	}

	dead := gscSubmit(p, "http://127.0.0.1:1/publish", "tok")
	if dead.Status != StatusFailed || !strings.Contains(string(dead.Response), `"status_code":0`) {
		t.Errorf("transport error = %+v, want failed with status_code 0", dead)
	}
}
