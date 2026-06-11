//ff:func feature=archive type=client control=iteration dimension=1
//ff:what processIndexNow가 일괄 1회 POST로 전 target done/failed를 함께 매기고, 인자 키가 env보다 우선하며, 키 부재·5xx면 전 항목 failed인지 검증
package archive

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessIndexNow(t *testing.T) {
	pending := []Pending{
		{Kind: KindIndexNow, Target: "https://blog.example.com/a/"},
		{Kind: KindIndexNow, Target: "https://blog.example.com/b/"},
	}
	calls := 0
	fail := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read body: %v", err)
		}
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil || payload["key"] != "k123" {
			t.Errorf("payload = %s (err=%v), want key k123", body, err)
		}
		if fail {
			http.Error(w, "indexnow down", http.StatusServiceUnavailable)
			return
		}
		if _, err := w.Write([]byte(`{}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("INDEXNOW_ENDPOINT", srv.URL)

	t.Setenv("INDEXNOW_KEY", "")
	for _, item := range processIndexNow("", pending) {
		if item.Status != StatusFailed || !strings.Contains(string(item.Response), "INDEXNOW_KEY") {
			t.Errorf("missing key: %+v, want failed with reason", item)
		}
	}
	if calls != 0 {
		t.Errorf("missing key must not call the endpoint (calls=%d)", calls)
	}

	t.Setenv("INDEXNOW_KEY", "k123")
	items := processIndexNow("", pending)
	if calls != 1 {
		t.Errorf("batch submission must POST exactly once, got %d", calls)
	}
	if len(items) != 2 || items[0].Status != StatusDone || items[1].Status != StatusDone {
		t.Errorf("2xx batch = %+v, want both done", items)
	}

	t.Setenv("INDEXNOW_KEY", "env-key-must-lose")
	for _, item := range processIndexNow("k123", pending) {
		if item.Status != StatusDone {
			t.Errorf("caller key must win over env: %+v, want done", item)
		}
	}

	fail = true
	for _, item := range processIndexNow("k123", pending) {
		if item.Status != StatusFailed {
			t.Errorf("5xx batch item = %+v, want failed", item)
		}
	}

	if processIndexNow("k123", nil) != nil {
		t.Error("empty group must be a no-op")
	}
}
