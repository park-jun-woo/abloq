//ff:func feature=cli type=command control=sequence
//ff:what claudemd 명령이 dir 인자를 받아 CLAUDE.md를 생성하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestNewClaudeMDCmd(t *testing.T) {
	dir := writeBlogFixture(t)
	cmd := newClaudeMDCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("claudemd: %v\noutput: %s", err, out.String())
	}
	if _, err := os.Stat(filepath.Join(dir, "CLAUDE.md")); err != nil {
		t.Errorf("CLAUDE.md missing: %v", err)
	}
}
