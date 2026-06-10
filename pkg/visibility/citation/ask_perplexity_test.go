//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what askPerplexity가 Bearer 키·모델·질의로 chat/completions를 호출하고 최상위 citations를 반환, 비2xx면 에러인지 검증
package citation

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAskPerplexity(t *testing.T) {
	badJSON := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/chat/completions" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if badJSON {
			_, _ = w.Write([]byte(`not-json`))
			return
		}
		if r.Header.Get("Authorization") != "Bearer px-key" {
			t.Errorf("authorization = %q", r.Header.Get("Authorization"))
		}
		body, _ := io.ReadAll(r.Body)
		var req struct {
			Model    string `json:"model"`
			Messages []struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"messages"`
		}
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("request body: %v", err)
		}
		if req.Model != "sonar" || len(req.Messages) != 1 || req.Messages[0].Content != "the question" {
			t.Errorf("request = %+v", req)
		}
		if req.Messages[0].Role != "user" {
			t.Errorf("role = %q", req.Messages[0].Role)
		}
		_, _ = w.Write([]byte(`{"citations":["https://blog.test/a/","https://other.example.org/"],"choices":[{"message":{"content":"answer"}}]}`))
	}))
	defer srv.Close()

	urls, err := askPerplexity(srv.URL, "px-key", "sonar", "the question")
	if err != nil {
		t.Fatalf("askPerplexity: %v", err)
	}
	want := []string{"https://blog.test/a/", "https://other.example.org/"}
	if !reflect.DeepEqual(urls, want) {
		t.Errorf("urls = %v, want %v", urls, want)
	}

	if _, err := askPerplexity("http://127.0.0.1:1", "px-key", "sonar", "q"); err == nil {
		t.Error("transport failure must fail")
	}
	badJSON = true
	if _, err := askPerplexity(srv.URL, "px-key", "sonar", "q"); err == nil {
		t.Error("malformed JSON must fail")
	}
}
