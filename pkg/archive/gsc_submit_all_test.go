//ff:func feature=archive type=client control=sequence
//ff:what gscSubmitAll이 target마다 publish하고 한 건의 실패가 나머지를 막지 않는지 검증
package archive

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGscSubmitAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read: %v", err)
		}
		var req map[string]string
		if err := json.Unmarshal(body, &req); err != nil || req["type"] != "URL_UPDATED" {
			t.Errorf("publish body = %s (err=%v)", body, err)
		}
		if strings.Contains(req["url"], "boom") {
			http.Error(w, "publish failed", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write([]byte(`{}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()

	items := gscSubmitAll([]Pending{
		{Kind: KindGSCIndex, Target: "https://blog/a/"},
		{Kind: KindGSCIndex, Target: "https://blog/boom/"},
	}, srv.URL, "tok")
	if len(items) != 2 {
		t.Fatalf("len = %d, want 2", len(items))
	}
	if items[0].Status != StatusDone || items[1].Status != StatusFailed {
		t.Errorf("statuses = %s/%s, want done/failed (isolated)", items[0].Status, items[1].Status)
	}
	if len(gscSubmitAll(nil, srv.URL, "tok")) != 0 {
		t.Error("empty group must produce no items")
	}
}
