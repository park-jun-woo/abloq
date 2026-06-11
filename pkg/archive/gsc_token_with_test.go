//ff:func feature=archive type=client control=sequence
//ff:what GSCTokenWith가 사이트 자격(인라인 JSON·파일 경로)으로 토큰을 받아오고 — 전역 env 없이 — 비2xx·빈 토큰·자격 불량·전송 실패·누락이면 에러인지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGSCTokenWith(t *testing.T) {
	fail := false
	empty := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fail {
			http.Error(w, "denied", http.StatusUnauthorized)
			return
		}
		if empty {
			if _, err := w.Write([]byte(`{}`)); err != nil {
				t.Errorf("write: %v", err)
			}
			return
		}
		if _, err := w.Write([]byte(`{"access_token":"stub-token","expires_in":3600}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("GSC_TOKEN_URL", srv.URL)
	// no global credentials: only the site-row Keys may make this pass
	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", "")

	saJSON := saJSONFixture(t)
	token, err := GSCTokenWith(Keys{GSCSAJSON: saJSON}, ScopeIndexing)
	if err != nil || token != "stub-token" {
		t.Errorf("inline site keys = %q, %v, want stub-token", token, err)
	}

	path := filepath.Join(t.TempDir(), "sa.json")
	if err := os.WriteFile(path, []byte(saJSON), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	token, err = GSCTokenWith(Keys{GSCSAJSONPath: path}, ScopeWebmastersReadonly)
	if err != nil || token != "stub-token" {
		t.Errorf("site sa_json_path = %q, %v, want stub-token", token, err)
	}

	fail = true
	if _, err := GSCTokenWith(Keys{GSCSAJSON: saJSON}, ScopeIndexing); err == nil {
		t.Error("non-2xx token endpoint must fail")
	}
	fail, empty = false, true
	if _, err := GSCTokenWith(Keys{GSCSAJSON: saJSON}, ScopeIndexing); err == nil {
		t.Error("empty access_token must fail")
	}

	fail, empty = false, false
	badKey := `{"client_email":"x@test","private_key":"not-pem"}`
	if _, err := GSCTokenWith(Keys{GSCSAJSON: badKey}, ScopeIndexing); err == nil {
		t.Error("unparsable private key must fail")
	}

	t.Setenv("GSC_TOKEN_URL", "http://127.0.0.1:1")
	if _, err := GSCTokenWith(Keys{GSCSAJSON: saJSON}, ScopeIndexing); err == nil {
		t.Error("transport failure must fail")
	}

	if _, err := GSCTokenWith(Keys{}, ScopeIndexing); err == nil {
		t.Error("empty keys without env credentials must fail")
	}
}
