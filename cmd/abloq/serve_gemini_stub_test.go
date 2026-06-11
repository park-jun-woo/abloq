//ff:func feature=cli type=command control=sequence
//ff:what 테스트 픽스처 — Gemini API 로컬 스텁 서버(GEMINI_API_BASE용), 모델 boom-model은 500·그 외는 inline base64 PNG 응답 (실 네트워크 0)
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func serveGeminiStub(t *testing.T) *httptest.Server {
	t.Helper()
	var buf bytes.Buffer
	if err := png.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 800, 800))); err != nil {
		t.Fatalf("png encode: %v", err)
	}
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "boom-model") {
			http.Error(w, `{"error":"boom-model is on fire"}`, http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, `{"candidates":[{"content":{"parts":[{"inlineData":{"mimeType":"image/png","data":%q}}]}}]}`, b64)
	}))
	t.Cleanup(srv.Close)
	return srv
}
