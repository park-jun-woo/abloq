//ff:func feature=cli type=command control=sequence topic=drift
//ff:what abloq check가 generate 직후 통과하고 수동 변조 시 파일·룰ID 진단과 함께 실패하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCheck(t *testing.T) {
	dir := writeBlogFixture(t)
	var out bytes.Buffer
	if err := runGenerate(&out, dir); err != nil {
		t.Fatalf("runGenerate: %v", err)
	}
	if err := runCheck(&out, dir); err != nil {
		t.Fatalf("runCheck after generate: %v\noutput: %s", err, out.String())
	}
	if !strings.Contains(out.String(), "in sync") {
		t.Errorf("want 'in sync' in output, got %q", out.String())
	}
	llmsPath := filepath.Join(dir, "static", "llms.txt")
	if err := os.WriteFile(llmsPath, []byte("# tampered\n"), 0o644); err != nil {
		t.Fatalf("tamper: %v", err)
	}
	out.Reset()
	err := runCheck(&out, dir)
	if err == nil {
		t.Fatalf("want error after tampering, got nil\noutput: %s", out.String())
	}
	if !strings.Contains(out.String(), "[llmstxt-sync]") {
		t.Errorf("want [llmstxt-sync] diagnostic, got %q", out.String())
	}
}
