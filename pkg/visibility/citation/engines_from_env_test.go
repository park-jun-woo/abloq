//ff:func feature=visibility type=client control=iteration dimension=1 topic=citation
//ff:what EnginesFromEnv가 키 있는 엔진만 perplexity·openai·anthropic 고정 순서로 조립하고 Ask가 env 베이스의 각 API를 실제 호출하는지 검증
package citation

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnginesFromEnv(t *testing.T) {
	t.Setenv("PERPLEXITY_API_KEY", "")
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("ANTHROPIC_API_KEY", "")
	if got := EnginesFromEnv(); len(got) != 0 {
		t.Errorf("no keys = %d engines, want 0", len(got))
	}

	t.Setenv("PERPLEXITY_API_KEY", "px")
	t.Setenv("ANTHROPIC_API_KEY", "an")
	engines := EnginesFromEnv()
	if len(engines) != 2 || engines[0].Name != "perplexity" || engines[1].Name != "anthropic" {
		names := make([]string, len(engines))
		for i, e := range engines {
			names[i] = e.Name
		}
		t.Fatalf("engines = %v, want [perplexity anthropic]", names)
	}

	t.Setenv("OPENAI_API_KEY", "oa")
	engines = EnginesFromEnv()
	if len(engines) != 3 || engines[1].Name != "openai" {
		t.Errorf("all keys = %d engines, second = %q", len(engines), engines[1].Name)
	}

	// The closures bind the env base URLs: every Ask must reach its own
	// endpoint on the stub.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/chat/completions":
			_, _ = w.Write([]byte(`{"citations":["https://blog.test/px/"]}`))
		case "/v1/responses":
			_, _ = w.Write([]byte(`{"output":[{"type":"message","content":[{"type":"output_text","annotations":[{"type":"url_citation","url":"https://blog.test/oa/"}]}]}]}`))
		case "/v1/messages":
			_, _ = w.Write([]byte(`{"content":[{"type":"text","citations":[{"type":"web_search_result_location","url":"https://blog.test/an/"}]}]}`))
		default:
			t.Errorf("unexpected path %q", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	t.Setenv("PERPLEXITY_BASE_URL", srv.URL)
	t.Setenv("OPENAI_BASE_URL", srv.URL)
	t.Setenv("ANTHROPIC_BASE_URL", srv.URL)
	want := map[string]string{
		"perplexity": "https://blog.test/px/",
		"openai":     "https://blog.test/oa/",
		"anthropic":  "https://blog.test/an/",
	}
	for _, e := range EnginesFromEnv() {
		urls, err := e.Ask("q")
		if err != nil || len(urls) != 1 || urls[0] != want[e.Name] {
			t.Errorf("%s.Ask = %v, %v, want [%s]", e.Name, urls, err, want[e.Name])
		}
	}
}
