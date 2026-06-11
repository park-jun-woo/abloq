//ff:func feature=cli type=command control=sequence topic=drift
//ff:what llms_txt: manual 선언 시 손작성 llms.txt가 generate에 덮어쓰이지 않고 check도 무변경 통과하는지 검증 (BUG001 수용 기준)
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCheckManualLlms(t *testing.T) {
	dir := writeBlogFixture(t)
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [opinion]\ngeo:\n  llms_txt: manual\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	curated := []byte("# Hand-curated llms.txt\n\n- [Master](/reins.md)\n")
	llmsPath := filepath.Join(dir, "static", "llms.txt")
	if err := os.MkdirAll(filepath.Dir(llmsPath), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(llmsPath, curated, 0o644); err != nil {
		t.Fatalf("write curated llms.txt: %v", err)
	}
	var out bytes.Buffer
	if err := runGenerate(&out, dir); err != nil {
		t.Fatalf("runGenerate: %v\noutput: %s", err, out.String())
	}
	got, err := os.ReadFile(llmsPath)
	if err != nil || !bytes.Equal(got, curated) {
		t.Errorf("generate must leave the curated llms.txt untouched (err %v), got %q", err, got)
	}
	if err := runCheck(&out, dir); err != nil {
		t.Fatalf("runCheck with curated llms.txt: %v\noutput: %s", err, out.String())
	}
	if strings.Contains(out.String(), "llmstxt-sync") {
		t.Errorf("manual mode must not enforce llmstxt-sync, got %q", out.String())
	}
}
