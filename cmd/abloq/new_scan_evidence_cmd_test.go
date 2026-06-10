//ff:func feature=cli type=command control=sequence
//ff:what evidence 명령이 Use/최대 1개 인자를 선언하고 RunE가 dir 인자를 실행 본체로 넘기는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewScanEvidenceCmd(t *testing.T) {
	cmd := newScanEvidenceCmd()
	if cmd.Use != "evidence [dir]" {
		t.Errorf("Use = %q, want \"evidence [dir]\"", cmd.Use)
	}
	if err := cmd.Args(cmd, []string{"a", "b"}); err == nil {
		t.Error("two args must be rejected")
	}
	if err := cmd.Args(cmd, []string{}); err != nil {
		t.Errorf("zero args rejected: %v", err)
	}
	// RunE wires the dir argument into the scan body (fixture has no claims).
	dir := writeBlogFixture(t)
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := cmd.RunE(cmd, []string{dir}); err != nil {
		t.Fatalf("RunE: %v", err)
	}
	if !strings.Contains(out.String(), "0 article(s) queued") {
		t.Errorf("scan summary missing: %q", out.String())
	}
	// Zero args default to the current directory — no blog.yaml there.
	if err := cmd.RunE(cmd, nil); err == nil {
		t.Error("default dir without blog.yaml must error")
	}
}
