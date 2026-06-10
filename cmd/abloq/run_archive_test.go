//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runArchive가 kind별 결과 3줄을 출력하고, 전부 done이면 성공·하나라도 실패면 에러인지 검증
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunArchive(t *testing.T) {
	gscOK := true
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			if !gscOK {
				http.Error(w, "denied", http.StatusUnauthorized)
				return
			}
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
	t.Setenv("GSC_SA_JSON", testSAJSONFixture(t))
	t.Setenv("GSC_SA_JSON_PATH", "")

	var out strings.Builder
	if err := runArchive(&out, "https://blog.example.com/p/"); err != nil {
		t.Fatalf("runArchive: %v\n%s", err, out.String())
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 3 {
		t.Errorf("output lines = %d, want 3 (one per kind)\n%s", len(lines), out.String())
	}
	for _, want := range []string{"wayback\tdone", "indexnow\tdone", "gsc_index\tdone"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("output lacks %q:\n%s", want, out.String())
		}
	}

	gscOK = false
	out.Reset()
	if err := runArchive(&out, "https://blog.example.com/p/"); err == nil {
		t.Error("a failed submission must surface as an error")
	}
}
