//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what verifyCitation 케이스 — 200+메타 ok, 404 broken, 5xx retry, 메타 불일치, 접속 불가 retry
package gate

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestVerifyCitation(t *testing.T) {
	srv := httptest.NewServer(newCitationMux())
	defer srv.Close()
	cases := []struct {
		name, url, label    string
		wantVerdict, detail string
	}{
		{"200 with matching title", srv.URL + "/ok", "Example Benchmark Report", "ok", ""},
		{"404 is broken", srv.URL + "/missing", "Anything", "broken", "HTTP 404"},
		{"5xx is retry", srv.URL + "/boom", "Anything", "retry", "HTTP 500"},
		{"meta mismatch", srv.URL + "/mismatch", "Example Benchmark Report", "meta-mismatch", "og:title"},
		{"unreachable host is retry", "http://127.0.0.1:1/dead", "Anything", "retry", "refused"},
		{"truncated body is retry", srv.URL + "/truncated", "Anything", "retry", "read:"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			verdict, detail := verifyCitation(srv.Client(), Citation{Label: tc.label, URL: tc.url})
			if verdict != tc.wantVerdict || !strings.Contains(detail, tc.detail) {
				t.Errorf("verifyCitation = (%q, %q), want (%q, contains %q)", verdict, detail, tc.wantVerdict, tc.detail)
			}
		})
	}
}
