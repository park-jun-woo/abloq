//ff:func feature=scan type=rule control=iteration dimension=1 topic=evidence
//ff:what Scan 3회전 — 1·2회는 claims 항목만(무상태 CLI와 동일), 죽은 인용은 연속 실패 3회째에만 rot 항목으로 확정
package evidence

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestScanThreeStrikes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "dead") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	root := writeRepoFixture(t, srv.URL+"/dead-1")
	b := testBlog(t)
	ck := &Checker{Client: srv.Client(), UserAgent: "abloqd-linkcheck"}

	var prev []Check
	for pass := 1; pass <= 3; pass++ {
		res := Scan(root, b, prev, ck)
		wantItems := 1
		if pass == 3 {
			wantItems = 2 // post-cite's dead link reached 3 consecutive failures
		}
		if len(res.Items) != wantItems {
			t.Fatalf("pass %d: items = %d, want %d (%+v)", pass, len(res.Items), wantItems, res.Items)
		}
		if res.Items[0].Slug != "post-claims" || res.Items[0].Payload["claims"] == "" {
			t.Errorf("pass %d: claims item = %+v", pass, res.Items[0])
		}
		if len(res.Checks) != 1 || res.Checks[0].ConsecutiveFailures != int64(pass) {
			t.Errorf("pass %d: checks = %+v, want 1 dead check at %d failures", pass, res.Checks, pass)
		}
		if pass == 3 && res.Items[1].Payload["rot_urls"] != `["`+srv.URL+`/dead-1"]` {
			t.Errorf("rot item payload = %+v", res.Items[1].Payload)
		}
		prev = res.Checks
	}
}
