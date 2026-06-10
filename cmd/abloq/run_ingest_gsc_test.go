//ff:func feature=cli type=command control=sequence topic=gsc
//ff:what runIngestGSC가 토큰 교환 후 닫힌 일자들을 조회해 행·합계를 출력하고, --site 결손은 blog.yaml baseURL로 보완하는지 검증
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunIngestGSC(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			_, _ = w.Write([]byte(`{"access_token":"stub-token"}`))
			return
		}
		if !strings.Contains(r.URL.EscapedPath(), "searchAnalytics/query") {
			t.Errorf("unexpected path %q", r.URL.EscapedPath())
		}
		if !strings.Contains(r.URL.EscapedPath(), "https:%2F%2Ft.example.com%2F") {
			t.Errorf("site not derived from blog.yaml: %q", r.URL.EscapedPath())
		}
		_, _ = w.Write([]byte(`{"rows":[{"keys":["https://t.example.com/opinion/hello/"],"clicks":2,"impressions":50,"position":3.1}]}`))
	}))
	defer srv.Close()
	t.Setenv("GSC_TOKEN_URL", srv.URL+"/token")
	t.Setenv("GSC_SA_JSON", testSAJSONFixture(t))
	t.Setenv("GSC_SA_JSON_PATH", "")
	t.Setenv("GSC_SEARCH_API_BASE", srv.URL)

	repo := writeBlogFixture(t)
	var out bytes.Buffer
	if err := runIngestGSC(&out, "", repo, 2); err != nil {
		t.Fatalf("runIngestGSC: %v", err)
	}
	got := out.String()
	if !strings.Contains(got, "https://t.example.com/opinion/hello/") {
		t.Errorf("row missing:\n%s", got)
	}
	if !strings.Contains(got, "gsc: 2 closed day(s), 2 row(s)") {
		t.Errorf("summary missing:\n%s", got)
	}

	// Search Analytics failure aborts the run with an error.
	t.Setenv("GSC_SEARCH_API_BASE", "http://127.0.0.1:1")
	if err := runIngestGSC(&out, "sc-domain:t.example.com", repo, 1); err == nil {
		t.Error("collect failure accepted")
	}

	if err := runIngestGSC(&out, "", t.TempDir(), 1); err == nil {
		t.Error("repo without blog.yaml accepted")
	}

	t.Setenv("GSC_SA_JSON", "")
	if err := runIngestGSC(&out, "sc-domain:t.example.com", repo, 1); err == nil {
		t.Error("missing credentials accepted")
	}
}
