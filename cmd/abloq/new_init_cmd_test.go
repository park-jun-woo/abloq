//ff:func feature=init type=command control=sequence
//ff:what init 명령이 플래그만으로(비대화형) 블로그를 스캐폴드하는지 검증 — 에이전트 경로
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewInitCmd(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "blog")
	cmd := newInitCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetIn(strings.NewReader(""))
	cmd.SetArgs([]string{dir, "--title", "Agent Blog", "--languages", "en", "--sections", "posts"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("init: %v\noutput: %s", err, out.String())
	}
	data, err := os.ReadFile(filepath.Join(dir, "blog.yaml"))
	if err != nil || !strings.Contains(string(data), `title: "Agent Blog"`) {
		t.Errorf("blog.yaml = %q, err %v", data, err)
	}
}
