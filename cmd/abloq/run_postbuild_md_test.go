//ff:func feature=cli type=command control=sequence
//ff:what runPostbuildMD가 생성 수를 출력하고 .md 본문이 front matter 없이 "# title"로 시작하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunPostbuildMD(t *testing.T) {
	dir := writeBlogFixture(t)
	var out bytes.Buffer
	if err := runPostbuildMD(&out, dir); err != nil {
		t.Fatalf("runPostbuildMD: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "public", "ko", "opinion", "hello.md"))
	if err != nil {
		t.Fatalf("read served md: %v", err)
	}
	got := string(data)
	if !strings.HasPrefix(got, "# Hello\n") || strings.Contains(got, "---") {
		t.Errorf("served md must be noise-free, got %q", got)
	}
	blocked := writeBlogFixture(t)
	if err := os.WriteFile(filepath.Join(blocked, "public"), []byte(""), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := runPostbuildMD(&out, blocked); err == nil {
		t.Error("public/ blocked by a regular file must error")
	}
}
