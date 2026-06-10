//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what postJSON이 JSON 헤더·Bearer 토큰으로 POST하고 2xx 본문 반환, 비2xx·전송 실패면 에러인지 검증
package gsc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" || r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("headers = %v", r.Header)
		}
		if r.URL.Path == "/fail" {
			http.Error(w, `{"error":"nope"}`, http.StatusBadRequest)
			return
		}
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	body, err := postJSON(srv.URL+"/ok", "tok", map[string]string{"a": "b"})
	if err != nil || string(body) != `{"ok":true}` {
		t.Errorf("postJSON = %q, %v", body, err)
	}

	if _, err := postJSON(srv.URL+"/fail", "tok", map[string]string{}); err == nil {
		t.Error("non-2xx must fail")
	}
	if _, err := postJSON("http://127.0.0.1:1", "tok", map[string]string{}); err == nil {
		t.Error("transport failure must fail")
	}
	if _, err := postJSON(srv.URL, "tok", func() {}); err == nil {
		t.Error("unmarshalable payload must fail")
	}
	if _, err := postJSON("http://bad url with spaces", "tok", map[string]string{}); err == nil {
		t.Error("invalid URL must fail")
	}
}
