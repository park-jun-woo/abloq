//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 인용 검증 테스트 공용 모의 서버 핸들러 — /ok 일치 제목, /redirect→/ok, /mismatch 무관 제목, /boom 500, /hang 지연
package gate

import (
	"net/http"
	"time"
)

func newCitationMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`<html><head><title>Example Benchmark Report</title></head></html>`))
	})
	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/ok", http.StatusFound)
	})
	mux.HandleFunc("/mismatch", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`<html><head><title>전혀 다른 페이지</title></head></html>`))
	})
	mux.HandleFunc("/boom", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	})
	mux.HandleFunc("/hang", func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(300 * time.Millisecond)
	})
	mux.HandleFunc("/truncated", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", "100")
		_, _ = w.Write([]byte("short"))
	})
	return mux
}
