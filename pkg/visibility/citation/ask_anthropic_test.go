//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what askAnthropic이 x-api-key·anthropic-version 헤더와 web_search 도구로 Messages API를 호출하고 content[].citations URL을 수집하는지 검증
package citation

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAskAnthropic(t *testing.T) {
	badJSON := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/messages" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if badJSON {
			_, _ = w.Write([]byte(`not-json`))
			return
		}
		if r.Header.Get("X-Api-Key") != "an-key" || r.Header.Get("Anthropic-Version") != "2023-06-01" {
			t.Errorf("headers = %v", r.Header)
		}
		body, _ := io.ReadAll(r.Body)
		var req struct {
			Model     string           `json:"model"`
			MaxTokens int              `json:"max_tokens"`
			Tools     []map[string]any `json:"tools"`
		}
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("request body: %v", err)
		}
		if req.Model != "claude-sonnet-4-5" || req.MaxTokens != 1024 ||
			len(req.Tools) != 1 || req.Tools[0]["type"] != "web_search_20250305" {
			t.Errorf("request = %+v", req)
		}
		_, _ = w.Write([]byte(`{"content":[
			{"type":"server_tool_use"},
			{"type":"text","text":"answer","citations":[
				{"type":"web_search_result_location","url":"https://blog.test/c/"},
				{"type":"web_search_result_location","url":"https://other.example.org/"}
			]}
		]}`))
	}))
	defer srv.Close()

	urls, err := askAnthropic(srv.URL, "an-key", "claude-sonnet-4-5", "the question")
	if err != nil {
		t.Fatalf("askAnthropic: %v", err)
	}
	want := []string{"https://blog.test/c/", "https://other.example.org/"}
	if !reflect.DeepEqual(urls, want) {
		t.Errorf("urls = %v, want %v", urls, want)
	}

	badJSON = true
	if _, err := askAnthropic(srv.URL, "an-key", "claude-sonnet-4-5", "q"); err == nil {
		t.Error("malformed JSON must fail")
	}
	if _, err := askAnthropic("http://127.0.0.1:1", "an-key", "claude-sonnet-4-5", "q"); err == nil {
		t.Error("transport failure must fail")
	}
}
