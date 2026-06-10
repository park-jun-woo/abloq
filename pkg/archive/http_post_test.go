//ff:func feature=archive type=client control=sequence
//ff:what httpPost가 상태코드·본문·헤더 전달을 보존하고, 연결 불가는 에러로 돌려주는지 검증
package archive

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPPost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("Authorization = %q, want Bearer tok", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusTeapot)
		if _, err := w.Write([]byte("short and stout")); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()

	header := http.Header{"Authorization": {"Bearer tok"}}
	code, body, err := httpPost(srv.URL, header, []byte("payload"))
	if err != nil {
		t.Fatalf("httpPost: %v", err)
	}
	if code != http.StatusTeapot || string(body) != "short and stout" {
		t.Errorf("got %d %q, want 418 \"short and stout\"", code, body)
	}

	if _, _, err := httpPost("http://127.0.0.1:1/unreachable", nil, nil); err == nil {
		t.Error("httpPost to a closed port must fail")
	}
	if _, _, err := httpPost("://bad", nil, nil); err == nil {
		t.Error("httpPost with an invalid URL must fail")
	}
}
