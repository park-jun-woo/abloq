//ff:func feature=cli type=command control=sequence
//ff:what runClaudeMD가 blog.yaml 값을 담은 CLAUDE.md를 기록하고 무효 blog.yaml에서 에러를 내는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunClaudeMD(t *testing.T) {
	dir := writeBlogFixture(t)
	var out bytes.Buffer
	if err := runClaudeMD(&out, dir); err != nil {
		t.Fatalf("runClaudeMD: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "CLAUDE.md"))
	if err != nil || !strings.Contains(string(data), "T 운영 매뉴얼") {
		t.Errorf("CLAUDE.md = %q, err %v", data, err)
	}
	if err := runClaudeMD(&out, t.TempDir()); err == nil {
		t.Error("missing blog.yaml must error")
	}
	blocked := writeBlogFixture(t)
	if err := os.MkdirAll(filepath.Join(blocked, "CLAUDE.md"), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := runClaudeMD(&out, blocked); err == nil {
		t.Error("CLAUDE.md blocked by a directory must error")
	}
}
