//ff:func feature=archive type=client control=iteration dimension=1
//ff:what processGSC가 신규 우선으로 쿼터까지 제출하고 초과분을 deferred로 이월하며, 토큰 실패면 제출분 전체 failed인지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessGSC(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			if _, err := w.Write([]byte(`{"access_token":"stub-token"}`)); err != nil {
				t.Errorf("write: %v", err)
			}
			return
		}
		if r.URL.Path != "/v3/urlNotifications:publish" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		if _, err := w.Write([]byte(`{}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("GSC_API_BASE", srv.URL)
	t.Setenv("GSC_TOKEN_URL", srv.URL+"/token")
	t.Setenv("GSC_SA_JSON", saJSONFixture(t))
	t.Setenv("GSC_SA_JSON_PATH", "")

	pending := []Pending{
		{Kind: KindGSCIndex, Target: "https://blog/updated/", Date: "2026-01-01", Lastmod: "2026-02-01"},
		{Kind: KindGSCIndex, Target: "https://blog/new/", Date: "2026-01-01", Lastmod: "2026-01-01"},
	}

	t.Setenv("GSC_DAILY_QUOTA", "1")
	byTarget := map[string]Item{}
	for _, item := range processGSC(Keys{}, pending) {
		byTarget[item.Target] = item
	}
	if len(byTarget) != 2 {
		t.Fatalf("items = %d, want 2", len(byTarget))
	}
	if byTarget["https://blog/new/"].Status != StatusDone {
		t.Errorf("new post = %s, want done (priority wins the quota)", byTarget["https://blog/new/"].Status)
	}
	deferredItem := byTarget["https://blog/updated/"]
	if deferredItem.Status != StatusDeferred || !strings.Contains(string(deferredItem.Response), "quota split") {
		t.Errorf("updated post = %+v, want deferred with reason", deferredItem)
	}

	t.Setenv("GSC_DAILY_QUOTA", "")
	t.Setenv("GSC_SA_JSON", "")
	for _, item := range processGSC(Keys{}, pending) {
		if item.Status != StatusFailed {
			t.Errorf("token failure item = %+v, want failed", item)
		}
	}

	if processGSC(Keys{}, nil) != nil {
		t.Error("empty group must be a no-op")
	}
}
