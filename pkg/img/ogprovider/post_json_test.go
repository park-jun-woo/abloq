//ff:func feature=image type=client control=sequence
//ff:what postJSON 검증 — 로컬 스텁으로 JSON POST·키 헤더·본문 반환, 비2xx의 본문 동봉 에러, 마샬 불가 페이로드 에러
package ogprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "wrong method: "+r.Method, http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "missing content type", http.StatusBadRequest)
			return
		}
		if r.Header.Get("x-goog-api-key") != "test-key" {
			http.Error(w, "missing api key header", http.StatusUnauthorized)
			return
		}
		body, _ := io.ReadAll(r.Body)
		var got map[string]string
		if err := json.Unmarshal(body, &got); err != nil || got["prompt"] != "p" {
			http.Error(w, "bad payload: "+string(body), http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer srv.Close()

	data, err := postJSON(context.Background(), srv.URL, "test-key", map[string]string{"prompt": "p"})
	if err != nil || string(data) != `{"ok":true}` {
		t.Fatalf("postJSON = %q, %v", data, err)
	}

	// non-2xx surfaces status + body snippet truncated to 512 bytes
	long := strings.Repeat("e", 600)
	fail := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error":"quota exceeded `+long+`"}`, http.StatusTooManyRequests)
	}))
	defer fail.Close()
	err = nil
	if _, err = postJSON(context.Background(), fail.URL, "k", map[string]string{}); err == nil ||
		!strings.Contains(err.Error(), "429") || !strings.Contains(err.Error(), "quota exceeded") {
		t.Errorf("non-2xx: want status + body snippet, got %v", err)
	}
	if err != nil && len(err.Error()) > 600 {
		t.Errorf("snippet not truncated: %d bytes", len(err.Error()))
	}

	// unmarshalable payload fails before any request
	if _, err := postJSON(context.Background(), srv.URL, "k", func() {}); err == nil {
		t.Error("unmarshalable payload must error")
	}

	// invalid endpoint fails at request construction
	if _, err := postJSON(context.Background(), "http://bad\x7fhost/", "k", map[string]string{}); err == nil {
		t.Error("invalid endpoint must error")
	}

	// unreachable endpoint fails at Do (closed local listener, no network)
	dead := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	deadURL := dead.URL
	dead.Close()
	if _, err := postJSON(context.Background(), deadURL, "k", map[string]string{}); err == nil {
		t.Error("closed endpoint must error")
	}

	// body read failure: declared length longer than the delivered body
	short := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "100")
		w.Write([]byte("partial"))
	}))
	defer short.Close()
	if _, err := postJSON(context.Background(), short.URL, "k", map[string]string{}); err == nil {
		t.Error("truncated body must error")
	}
}
