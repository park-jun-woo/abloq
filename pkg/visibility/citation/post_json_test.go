//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what postJSON이 호출자 헤더 위에 JSON 헤더를 얹어 POST하고 2xx 본문 반환, 비2xx·전송 실패면 에러인지 검증
package citation

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostJSONCitation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" || r.Header.Get("X-Api-Key") != "k" {
			t.Errorf("headers = %v", r.Header)
		}
		if r.URL.Path == "/fail" {
			http.Error(w, `{"error":"nope"}`, http.StatusTooManyRequests)
			return
		}
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	body, err := postJSON(srv.URL+"/ok", http.Header{"X-Api-Key": {"k"}}, map[string]string{"a": "b"})
	if err != nil || string(body) != `{"ok":true}` {
		t.Errorf("postJSON = %q, %v", body, err)
	}

	if _, err := postJSON(srv.URL+"/fail", http.Header{"X-Api-Key": {"k"}}, map[string]string{}); err == nil {
		t.Error("non-2xx must fail")
	}
	if _, err := postJSON("http://127.0.0.1:1", nil, map[string]string{}); err == nil {
		t.Error("transport failure must fail")
	}
	if _, err := postJSON(srv.URL, nil, func() {}); err == nil {
		t.Error("unmarshalable payload must fail")
	}
	if _, err := postJSON("http://bad url with spaces", nil, map[string]string{}); err == nil {
		t.Error("invalid URL must fail")
	}
}
