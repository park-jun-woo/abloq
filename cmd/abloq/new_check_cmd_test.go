//ff:func feature=cli type=command control=sequence topic=drift
//ff:what check 명령이 dir 인자를 받아 실행되고 파생물 누락 시 exit용 에러를 반환하는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewCheckCmd(t *testing.T) {
	dir := writeBlogFixture(t)
	cmd := newCheckCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("want error before generate (derived files missing), got nil\noutput: %s", out.String())
	}
	if !strings.Contains(out.String(), "missing") {
		t.Errorf("want missing-file diagnostics, got %q", out.String())
	}
}
