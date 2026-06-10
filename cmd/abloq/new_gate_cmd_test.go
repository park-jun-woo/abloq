//ff:func feature=cli type=command control=sequence
//ff:what gate 명령이 dir 인자를 받아 실행되고 위반 글에서 룰ID 진단과 에러를 반환하는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewGateCmd(t *testing.T) {
	dir := writeGateFixture(t, false)
	cmd := newGateCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("want error for violating articles, got nil\noutput: %s", out.String())
	}
	if !strings.Contains(out.String(), "[image-first]") {
		t.Errorf("want [image-first] diagnostic, got %q", out.String())
	}
}
