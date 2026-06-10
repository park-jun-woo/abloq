//ff:func feature=gen type=rule control=sequence topic=drift
//ff:what Check가 동기 상태에서 0 진단, 파일 누락·수동 변조에서 룰ID 달린 진단을 내는지 검증
package gen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheck(t *testing.T) {
	dir := t.TempDir()
	outs := []Output{
		{Path: "static/robots.txt", Data: []byte("User-agent: GPTBot\nDisallow:\n")},
		{Path: "static/llms.txt", Data: []byte("# X\n")},
	}
	if err := Write(dir, outs); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if diags := Check(dir, outs); len(diags) != 0 {
		t.Fatalf("in-sync Check: want 0 diagnostics, got %v", diags)
	}
	tampered := filepath.Join(dir, "static", "robots.txt")
	if err := os.WriteFile(tampered, []byte("User-agent: GPTBot\nDisallow: /\n"), 0o644); err != nil {
		t.Fatalf("tamper: %v", err)
	}
	if err := os.Remove(filepath.Join(dir, "static", "llms.txt")); err != nil {
		t.Fatalf("remove: %v", err)
	}
	diags := Check(dir, outs)
	if len(diags) != 2 {
		t.Fatalf("want 2 diagnostics, got %v", diags)
	}
	checkDriftDiagFields(t, diags[0], tampered, 2, "robots-policy-match", `want "Disallow:", got "Disallow: /"`)
	if diags[1].Rule != "llmstxt-sync" || diags[1].Line != 1 {
		t.Errorf("missing-file diagnostic = %+v, want rule llmstxt-sync line 1", diags[1])
	}
}
