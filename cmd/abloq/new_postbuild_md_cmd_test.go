//ff:func feature=cli type=command control=sequence
//ff:what postbuild md 명령이 dir 인자를 받아 글별 .md를 산출하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewPostbuildMDCmd(t *testing.T) {
	dir := writeBlogFixture(t)
	cmd := newPostbuildMDCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("postbuild md: %v\noutput: %s", err, out.String())
	}
	if !strings.Contains(out.String(), "postbuild md: 1 file(s)") {
		t.Errorf("want count output, got %q", out.String())
	}
	if _, err := os.Stat(filepath.Join(dir, "public", "ko", "opinion", "hello.md")); err != nil {
		t.Errorf("served md missing: %v", err)
	}
}
