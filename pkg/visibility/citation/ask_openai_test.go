//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what askOpenAI가 web_search 도구로 Responses API를 호출하고 url_citation 어노테이션 URL만 수집하는지 검증
package citation

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAskOpenAI(t *testing.T) {
	badJSON := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/responses" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if badJSON {
			_, _ = w.Write([]byte(`not-json`))
			return
		}
		if r.Header.Get("Authorization") != "Bearer oa-key" {
			t.Errorf("authorization = %q", r.Header.Get("Authorization"))
		}
		body, _ := io.ReadAll(r.Body)
		var req struct {
			Model string              `json:"model"`
			Tools []map[string]string `json:"tools"`
			Input string              `json:"input"`
		}
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("request body: %v", err)
		}
		if req.Model != "gpt-4.1" || req.Input != "the question" ||
			len(req.Tools) != 1 || req.Tools[0]["type"] != "web_search" {
			t.Errorf("request = %+v", req)
		}
		_, _ = w.Write([]byte(`{"output":[
			{"type":"web_search_call","content":null},
			{"type":"message","content":[{"type":"output_text","text":"answer","annotations":[
				{"type":"url_citation","url":"https://blog.test/b/"},
				{"type":"file_citation","url":"https://skip.example.org/"},
				{"type":"url_citation","url":""}
			]}]}
		]}`))
	}))
	defer srv.Close()

	urls, err := askOpenAI(srv.URL, "oa-key", "gpt-4.1", "the question")
	if err != nil {
		t.Fatalf("askOpenAI: %v", err)
	}
	if !reflect.DeepEqual(urls, []string{"https://blog.test/b/"}) {
		t.Errorf("urls = %v, want only the url_citation", urls)
	}

	badJSON = true
	if _, err := askOpenAI(srv.URL, "oa-key", "gpt-4.1", "q"); err == nil {
		t.Error("malformed JSON must fail")
	}
	if _, err := askOpenAI("http://127.0.0.1:1", "oa-key", "gpt-4.1", "q"); err == nil {
		t.Error("transport failure must fail")
	}
}
