//ff:func feature=archive type=client control=iteration dimension=1
//ff:what ProcessBatchWith가 사이트 자격(Keys)을 indexnow/gsc 클라이언트로 흘리고(전역 env 없이 done) limit 상한을 지키는지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessBatchWith(t *testing.T) {
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
	t.Setenv("GSC_API_BASE", srv.URL)
	t.Setenv("GSC_TOKEN_URL", srv.URL+"/token")
	t.Setenv("GSC_DAILY_QUOTA", "")
	// no global credentials: only the site-row Keys may make this pass
	t.Setenv("INDEXNOW_KEY", "")
	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", "")

	keys := Keys{IndexNowKey: "site-key", GSCSAJSON: saJSONFixture(t)}
	pending := []Pending{
		{DeployID: "d", Kind: KindWayback, Target: "https://blog.example.com/p/"},
		{DeployID: "d", Kind: KindIndexNow, Target: "https://blog.example.com/p/"},
		{DeployID: "d", Kind: KindGSCIndex, Target: "https://blog.example.com/p/"},
	}
	results := ProcessBatchWith(keys, pending, 100)
	if len(results) != 3 {
		t.Fatalf("results = %d, want 3", len(results))
	}
	for _, item := range results {
		if item.Status != StatusDone {
			t.Errorf("%s = %s, want done (site keys must reach the clients)", item.Kind, item.Status)
		}
	}

	if got := ProcessBatchWith(keys, pending, 1); len(got) != 1 {
		t.Errorf("limit 1 processed %d, want 1 (rest stays pending)", len(got))
	}
	if got := ProcessBatchWith(keys, nil, 10); len(got) != 0 {
		t.Errorf("empty batch processed %d, want 0", len(got))
	}
}
