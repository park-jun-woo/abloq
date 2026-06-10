//ff:func feature=cli type=command control=sequence topic=citation
//ff:what runSampleCitations가 질의 파일·blog.yaml budget·env 키 엔진으로 1회전 실행해 샘플과 합계를 출력하고, budget 0이면 no-op 합계인지 검증
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunSampleCitations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"citations":["https://t.example.com/opinion/hello/"],"choices":[{"message":{"content":"a"}}]}`))
	}))
	defer srv.Close()
	t.Setenv("PERPLEXITY_API_KEY", "px")
	t.Setenv("PERPLEXITY_BASE_URL", srv.URL)
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("ANTHROPIC_API_KEY", "")
	t.Setenv("CITATION_INTERVAL_MS", "0")

	dir := t.TempDir()
	queries := filepath.Join(dir, "queries.yaml")
	if err := os.WriteFile(queries, []byte("- id: 1\n  query_text: q-one\n- id: 2\n  query_text: q-two\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	repo := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [opinion]\ngeo:\n  citation_budget: 1\n"
	if err := os.WriteFile(filepath.Join(repo, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	if err := runSampleCitations(&out, queries, repo); err != nil {
		t.Fatalf("runSampleCitations: %v", err)
	}
	got := out.String()
	if !strings.Contains(got, "perplexity      1 true") {
		t.Errorf("sample row missing:\n%s", got)
	}
	if !strings.Contains(got, "sample: 1 engine(s), budget 1, 1 sample(s) [extractor v1]") {
		t.Errorf("summary missing:\n%s", got)
	}

	// budget 0 (default) — sampling disabled, an explicit no-op.
	repo0 := writeBlogFixture(t)
	out.Reset()
	if err := runSampleCitations(&out, queries, repo0); err != nil {
		t.Fatalf("budget 0: %v", err)
	}
	if !strings.Contains(out.String(), "sample: 1 engine(s), budget 0, 0 sample(s)") {
		t.Errorf("budget-0 summary missing:\n%s", out.String())
	}

	if err := runSampleCitations(&out, filepath.Join(dir, "missing.yaml"), repo); err == nil {
		t.Error("missing queries file accepted")
	}
	if err := runSampleCitations(&out, queries, t.TempDir()); err == nil {
		t.Error("repo without blog.yaml accepted")
	}
}
