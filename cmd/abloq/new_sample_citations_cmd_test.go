//ff:func feature=cli type=command control=sequence topic=citation
//ff:what citations 명령이 --queries 필수·--repo 기본값을 선언하고 RunE가 실행 본체로 플래그를 넘기는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewSampleCitationsCmd(t *testing.T) {
	cmd := newSampleCitationsCmd()
	if !strings.HasPrefix(cmd.Use, "citations") {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.Flags().Lookup("queries") == nil || cmd.Flags().Lookup("repo") == nil {
		t.Fatal("--queries/--repo flags missing")
	}
	if got := cmd.Flags().Lookup("repo").DefValue; got != "." {
		t.Errorf("--repo default = %q, want \".\"", got)
	}

	// RunE wires the flags into the sampling body — no engine keys, budget
	// 0 fixture: the round is a clean no-op.
	t.Setenv("PERPLEXITY_API_KEY", "")
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("ANTHROPIC_API_KEY", "")
	dir := t.TempDir()
	queries := filepath.Join(dir, "queries.yaml")
	if err := os.WriteFile(queries, []byte("- id: 1\n  query_text: q\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	repo := writeBlogFixture(t)
	if err := cmd.Flags().Set("queries", queries); err != nil {
		t.Fatal(err)
	}
	if err := cmd.Flags().Set("repo", repo); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := cmd.RunE(cmd, nil); err != nil {
		t.Fatalf("RunE: %v", err)
	}
	if !strings.Contains(out.String(), "sample: 0 engine(s), budget 0, 0 sample(s)") {
		t.Errorf("summary missing: %q", out.String())
	}
}
