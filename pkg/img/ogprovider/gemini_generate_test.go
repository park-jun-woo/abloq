//ff:func feature=image type=client control=sequence
//ff:what Generate가 로컬 스텁 서버(GEMINI_API_BASE)로 inline base64 이미지를 받아 디코드하는지 — 경로/키 헤더/비2xx/이미지 부재 검증
package ogprovider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGeminiGenerate(t *testing.T) {
	b64 := stubPNGBase64(t, 640, 480)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1beta/models/test-model:generateContent" {
			http.Error(w, "wrong path: "+r.URL.Path, http.StatusNotFound)
			return
		}
		if r.Header.Get("x-goog-api-key") != "test-key" {
			http.Error(w, "missing api key header", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, `{"candidates":[{"content":{"parts":[{"text":"here"},{"inlineData":{"mimeType":"image/png","data":%q}}]}}]}`, b64)
	}))
	defer srv.Close()
	t.Setenv("GEMINI_API_KEY", "test-key")
	t.Setenv("GEMINI_API_BASE", srv.URL)

	g, err := NewGemini("test-model")
	if err != nil {
		t.Fatalf("NewGemini: %v", err)
	}
	m, err := g.Generate(context.Background(), "a prompt")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 640 || b.Dy() != 480 {
		t.Errorf("bounds = %v, want 640x480", b)
	}

	// non-2xx surfaces status + body snippet
	fail := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error":"quota exceeded"}`, http.StatusTooManyRequests)
	}))
	defer fail.Close()
	t.Setenv("GEMINI_API_BASE", fail.URL)
	g, _ = NewGemini("test-model")
	if _, err := g.Generate(context.Background(), "p"); err == nil || !strings.Contains(err.Error(), "quota") {
		t.Errorf("quota failure: want body snippet in error, got %v", err)
	}

	// response without an image part
	noimg := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"candidates":[{"content":{"parts":[{"text":"sorry"}]}}]}`)
	}))
	defer noimg.Close()
	t.Setenv("GEMINI_API_BASE", noimg.URL)
	g, _ = NewGemini("test-model")
	if _, err := g.Generate(context.Background(), "p"); err == nil || !strings.Contains(err.Error(), "no inline image") {
		t.Errorf("imageless response: want clear error, got %v", err)
	}
}
