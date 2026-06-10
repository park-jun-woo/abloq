//ff:func feature=archive type=client control=iteration dimension=1
//ff:what ProcessBatch가 limit 상한 적용 후 3종 클라이언트를 한 번에 실행하고 per-target 결과를 합치는지 검증 (스텁 4경로)
package archive

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessBatch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			if _, err := w.Write([]byte(`{"access_token":"stub-token"}`)); err != nil {
				t.Errorf("write: %v", err)
			}
			return
		}
		if _, err := w.Write([]byte(`{"ok":true}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("WAYBACK_BASE_URL", srv.URL)
	t.Setenv("INDEXNOW_ENDPOINT", srv.URL+"/indexnow")
	t.Setenv("INDEXNOW_KEY", "k123")
	t.Setenv("GSC_API_BASE", srv.URL)
	t.Setenv("GSC_TOKEN_URL", srv.URL+"/token")
	t.Setenv("GSC_SA_JSON", saJSONFixture(t))
	t.Setenv("GSC_SA_JSON_PATH", "")
	t.Setenv("GSC_DAILY_QUOTA", "")

	pending := []Pending{
		{DeployID: "d", Kind: KindWayback, Target: "https://blog.example.com/p/"},
		{DeployID: "d", Kind: KindIndexNow, Target: "https://blog.example.com/p/"},
		{DeployID: "d", Kind: KindGSCIndex, Target: "https://blog.example.com/p/"},
	}
	results := ProcessBatch(pending, 100)
	if len(results) != 3 {
		t.Fatalf("results = %d, want 3", len(results))
	}
	for _, item := range results {
		if item.Status != StatusDone {
			t.Errorf("%s = %s, want done", item.Kind, item.Status)
		}
	}

	if got := ProcessBatch(pending, 1); len(got) != 1 {
		t.Errorf("limit 1 processed %d, want 1 (rest stays pending)", len(got))
	}
	if got := ProcessBatch(pending, 0); len(got) != 0 {
		t.Errorf("limit 0 processed %d, want 0", len(got))
	}
	if got := ProcessBatch(nil, 10); len(got) != 0 {
		t.Errorf("empty batch processed %d, want 0", len(got))
	}
}
