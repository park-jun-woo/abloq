//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what [citation-exists] 모의 서버 5경로 — 200 PASS, 404 FAIL, 리다이렉트 PASS, 메타 불일치 FAIL, timeout은 RETRY 진단
package gate

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRuleCitationExists(t *testing.T) {
	srv := httptest.NewServer(newCitationMux())
	defer srv.Close()
	old := citationHTTP
	citationHTTP = &http.Client{Timeout: 100 * time.Millisecond} // /hang sleeps past this
	defer func() { citationHTTP = old }()
	cases := []struct {
		name, path, label string
		wantDiags         int
		wantMsgPart       string
	}{
		{"200 with matching meta passes", "/ok", "Example Benchmark Report", 0, ""},
		{"404 fails", "/missing", "Anything Else", 1, "HTTP 404"},
		{"redirect to 200 passes", "/redirect", "Example Benchmark Report", 0, ""},
		{"meta mismatch fails", "/mismatch", "Example Benchmark Report", 1, "og:title"},
		{"timeout is a RETRY diagnostic", "/hang", "Anything Else", 1, "RETRY"},
	}
	b := loadGateBlog(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := artFromContent(t, b, "Claim per ["+tc.label+"]("+srv.URL+tc.path+").\n")
			tgt := NewTarget(t.TempDir(), b, []*Article{a})
			checkDiags(t, ruleCitationExists(tgt), tc.wantDiags, "citation-exists", tc.wantMsgPart)
		})
	}
}
