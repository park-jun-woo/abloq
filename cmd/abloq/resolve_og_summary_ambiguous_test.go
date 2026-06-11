//ff:func feature=cli type=command control=sequence
//ff:what resolveOGSummary 모호 케이스 — base 언어에서 같은 slug로 2건이 겹치면 에러 아닌 빈값 폴백 + "모호" 진단 1줄
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolveOGSummaryAmbiguous(t *testing.T) {
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [en]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	enDir := filepath.Join(dir, "content", "en", "tech")
	if err := os.MkdirAll(enDir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	// two base-language articles collapse onto slug "dup" (one stem, one override)
	stem := "---\ntitle: A\ndate: 2026-01-02\nsummary: first\n---\nbody\n"
	override := "---\ntitle: B\ndate: 2026-01-03\nslug: dup\nsummary: second\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(enDir, "dup.md"), []byte(stem), 0o644); err != nil {
		t.Fatalf("write stem: %v", err)
	}
	if err := os.WriteFile(filepath.Join(enDir, "other.md"), []byte(override), 0o644); err != nil {
		t.Fatalf("write override: %v", err)
	}
	var out bytes.Buffer
	b, _, err := loadImageOG(&out, dir)
	if err != nil || b == nil {
		t.Fatalf("loadImageOG: %v", err)
	}
	out.Reset()
	if got := resolveOGSummary(&out, dir, b, "dup"); got != "" {
		t.Errorf("ambiguous slug = %q, want empty fallback", got)
	}
	if !strings.Contains(out.String(), "모호") {
		t.Errorf("ambiguous must print a diagnostic, got %q", out.String())
	}
}
