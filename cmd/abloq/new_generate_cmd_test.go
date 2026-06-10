//ff:func feature=cli type=command control=sequence
//ff:what generate 명령이 dir 인자를 받아 실행되고 기본 dir에 blog.yaml이 없으면 실패하는지 검증
package main

import (
	"bytes"
	"testing"
)

func TestNewGenerateCmd(t *testing.T) {
	dir := writeBlogFixture(t)
	cmd := newGenerateCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute: %v\noutput: %s", err, out.String())
	}
	missing := newGenerateCmd()
	missing.SetOut(&out)
	missing.SetErr(&out)
	missing.SetArgs([]string{t.TempDir()})
	if err := missing.Execute(); err == nil {
		t.Errorf("want error for dir without blog.yaml, got nil")
	}
}
