//ff:func feature=cli type=command control=sequence topic=gsc
//ff:what gsc 명령이 --site/--days/--repo 플래그와 기본값을 선언하고 RunE가 실행 본체로 플래그를 넘기는지 검증
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewIngestGSCCmd(t *testing.T) {
	cmd := newIngestGSCCmd()
	if !strings.HasPrefix(cmd.Use, "gsc") {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.Flags().Lookup("site") == nil || cmd.Flags().Lookup("days") == nil || cmd.Flags().Lookup("repo") == nil {
		t.Fatal("--site/--days/--repo flags missing")
	}
	if got := cmd.Flags().Lookup("days").DefValue; got != "7" {
		t.Errorf("--days default = %q, want \"7\"", got)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			_, _ = w.Write([]byte(`{"access_token":"stub-token"}`))
			return
		}
		_, _ = w.Write([]byte(`{"rows":[]}`))
	}))
	defer srv.Close()
	t.Setenv("GSC_TOKEN_URL", srv.URL+"/token")
	t.Setenv("GSC_SA_JSON", testSAJSONFixture(t))
	t.Setenv("GSC_SA_JSON_PATH", "")
	t.Setenv("GSC_SEARCH_API_BASE", srv.URL)

	if err := cmd.Flags().Set("site", "sc-domain:t.example.com"); err != nil {
		t.Fatal(err)
	}
	if err := cmd.Flags().Set("days", "1"); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := cmd.RunE(cmd, nil); err != nil {
		t.Fatalf("RunE: %v", err)
	}
	if !strings.Contains(out.String(), "gsc: 1 closed day(s), 0 row(s)") {
		t.Errorf("summary missing: %q", out.String())
	}
}
