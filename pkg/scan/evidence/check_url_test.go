//ff:func feature=scan type=client control=iteration dimension=1 topic=evidence
//ff:what checkURL 케이스 — 200 ok, 404 hard, 500 soft, HEAD 405면 GET 폴백, 접속 불가는 soft
package evidence

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/dead":
			w.WriteHeader(http.StatusNotFound)
		case "/flaky":
			w.WriteHeader(http.StatusInternalServerError)
		case "/no-head":
			if r.Method == http.MethodHead {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer srv.Close()
	c := &Checker{Client: srv.Client(), UserAgent: "abloqd-linkcheck"}
	cases := []struct{ path, want string }{
		{"/alive", "ok"}, {"/dead", "hard"}, {"/flaky", "soft"}, {"/no-head", "ok"},
	}
	for _, tc := range cases {
		if got := c.checkURL(srv.URL + tc.path); got != tc.want {
			t.Errorf("checkURL(%s) = %q, want %q", tc.path, got, tc.want)
		}
	}
	srv.Close()
	if got := c.checkURL(srv.URL + "/alive"); got != "soft" {
		t.Errorf("unreachable server = %q, want soft", got)
	}
	// HEAD is rejected as a method, then the GET fallback dies mid-flight.
	drop := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		hj, _ := w.(http.Hijacker)
		conn, _, err := hj.Hijack()
		if err != nil {
			t.Fatalf("hijack: %v", err)
		}
		conn.Close()
	}))
	defer drop.Close()
	dc := &Checker{Client: drop.Client(), UserAgent: "abloqd-linkcheck"}
	if got := dc.checkURL(drop.URL + "/x"); got != "soft" {
		t.Errorf("GET fallback transport error = %q, want soft", got)
	}
}
