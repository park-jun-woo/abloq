//ff:func feature=archive type=client control=sequence
//ff:what GSCToken이 jwt-bearer grant로 scope별 access_token을 받아오고 비2xx·빈 토큰·자격 누락이면 에러인지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGSCToken(t *testing.T) {
	fail := false
	empty := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		if r.PostForm.Get("grant_type") != "urn:ietf:params:oauth:grant-type:jwt-bearer" {
			t.Errorf("grant_type = %q", r.PostForm.Get("grant_type"))
		}
		if !strings.Contains(r.PostForm.Get("assertion"), ".") {
			t.Errorf("assertion %q is not a JWT", r.PostForm.Get("assertion"))
		}
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
	t.Setenv("GSC_SA_JSON", saJSONFixture(t))
	t.Setenv("GSC_SA_JSON_PATH", "")

	token, err := GSCToken(ScopeIndexing)
	if err != nil || token != "stub-token" {
		t.Errorf("GSCToken = %q, %v, want stub-token", token, err)
	}

	fail = true
	if _, err := GSCToken(ScopeIndexing); err == nil {
		t.Error("non-2xx token endpoint must fail")
	}
	fail, empty = false, true
	if _, err := GSCToken(ScopeIndexing); err == nil {
		t.Error("empty access_token must fail")
	}

	fail, empty = false, false
	t.Setenv("GSC_SA_JSON", `{"client_email":"x@test","private_key":"not-pem"}`)
	if _, err := GSCToken(ScopeIndexing); err == nil {
		t.Error("unparsable private key must fail")
	}

	t.Setenv("GSC_SA_JSON", saJSONFixture(t))
	t.Setenv("GSC_TOKEN_URL", "http://127.0.0.1:1")
	if _, err := GSCToken(ScopeIndexing); err == nil {
		t.Error("transport failure must fail")
	}

	t.Setenv("GSC_SA_JSON", "")
	if _, err := GSCToken(ScopeIndexing); err == nil {
		t.Error("missing credentials must fail")
	}
}
