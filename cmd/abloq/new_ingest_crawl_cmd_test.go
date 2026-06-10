//ff:func feature=cli type=command control=sequence topic=crawl
//ff:what crawl 명령이 --source 필수·--repo 기본값을 선언하고 RunE가 실행 본체로 플래그를 넘기는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewIngestCrawlCmd(t *testing.T) {
	cmd := newIngestCrawlCmd()
	if !strings.HasPrefix(cmd.Use, "crawl") {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.Flags().Lookup("source") == nil || cmd.Flags().Lookup("repo") == nil {
		t.Fatal("--source/--repo flags missing")
	}
	if got := cmd.Flags().Lookup("repo").DefValue; got != "." {
		t.Errorf("--repo default = %q, want \".\"", got)
	}
	// RunE wires the flags into the ingest body.
	repo := writeBlogFixture(t)
	logs := writeCFLogFixture(t)
	if err := cmd.Flags().Set("source", logs); err != nil {
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
	if !strings.Contains(out.String(), "1 file(s)") {
		t.Errorf("summary missing: %q", out.String())
	}
}
