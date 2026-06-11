package visibility

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// writeCitationRepo builds a blog repository whose blog.yaml carries the
// given citation_budget.
func writeCitationRepo(t *testing.T, budget string) string {
	t.Helper()
	repo := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [tech]\ngeo:\n  citation_budget: " + budget + "\n"
	if err := os.WriteFile(filepath.Join(repo, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	return repo
}

// TestSampleCitations runs the wrapper against an engine stub: every
// API-key engine answers the budgeted queries, the own-domain citation
// marks cited=true, and budget 0 is a clean no-op.
func TestSampleCitations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/chat/completions":
			w.Write([]byte(`{"citations":["https://t.example.com/tech/post-a/"],"choices":[{"message":{"content":"a"}}]}`))
		case "/v1/messages":
			w.Write([]byte(`{"content":[{"type":"text","text":"a","citations":[{"type":"web_search_result_location","url":"https://other.example.org/"}]}]}`))
		default:
			t.Errorf("unexpected path %q", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	t.Setenv("PERPLEXITY_API_KEY", "px")
	t.Setenv("PERPLEXITY_BASE_URL", srv.URL)
	t.Setenv("ANTHROPIC_API_KEY", "an")
	t.Setenv("ANTHROPIC_BASE_URL", srv.URL)
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("CITATION_INTERVAL_MS", "0")
	repo := writeCitationRepo(t, "1")

	queries := `[{"id":1,"query_text":"q-one"},{"id":2,"query_text":"q-two"}]`
	res, err := SampleCitations(SampleCitationsRequest{RepoPath: repo, QueriesJSON: queries})
	if err != nil {
		t.Fatalf("SampleCitations: %v", err)
	}
	if res.Engines != 2 || res.Sampled != 2 {
		t.Errorf("engines=%d sampled=%d, want 2/2 (budget 1 x 2 engines)", res.Engines, res.Sampled)
	}
	var samples []map[string]any
	if err := json.Unmarshal(res.SamplesJSON, &samples); err != nil {
		t.Fatalf("samples json: %v", err)
	}
	if len(samples) != 2 {
		t.Fatalf("samples = %d", len(samples))
	}
	if samples[0]["engine"] != "perplexity" || samples[0]["cited"] != true ||
		samples[0]["extractor_version"] != "v1" || samples[0]["citation_queries_id"] != float64(1) {
		t.Errorf("perplexity sample = %v", samples[0])
	}
	if samples[1]["engine"] != "anthropic" || samples[1]["cited"] != false {
		t.Errorf("anthropic sample = %v (other-domain citation must not count)", samples[1])
	}

	// budget 0 — sampling disabled, '[]' payload.
	res, err = SampleCitations(SampleCitationsRequest{RepoPath: writeCitationRepo(t, "0"), QueriesJSON: queries})
	if err != nil {
		t.Fatalf("budget 0: %v", err)
	}
	if res.Sampled != 0 || string(res.SamplesJSON) != "[]" {
		t.Errorf("budget-0 run = %+v", res)
	}

	if _, err := SampleCitations(SampleCitationsRequest{RepoPath: repo, QueriesJSON: "not-json"}); err == nil {
		t.Error("malformed queries JSON accepted")
	}
	badRepo := t.TempDir()
	badBlog := "site:\n  baseURL: not-a-url\n  title: T\n  author: A\nlanguages: [ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(badRepo, "blog.yaml"), []byte(badBlog), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := SampleCitations(SampleCitationsRequest{RepoPath: badRepo, QueriesJSON: "[]"}); err == nil {
		t.Error("invalid blog.yaml accepted")
	}
	if _, err := SampleCitations(SampleCitationsRequest{RepoPath: filepath.Join(badRepo, "missing"), QueriesJSON: "[]"}); err == nil {
		t.Error("missing blog.yaml accepted")
	}
	if _, err := SampleCitations(SampleCitationsRequest{QueriesJSON: "[]"}); err == nil {
		t.Error("missing repo_path accepted")
	}
}
