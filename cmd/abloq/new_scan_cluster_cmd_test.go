//ff:func feature=cli type=command control=sequence
//ff:what cluster 명령이 Use/최대 1개 인자를 선언하고 dir 인자 유무 양쪽에서 실행되는지 검증
package main

import (
	"bytes"
	"testing"
)

func TestNewScanClusterCmd(t *testing.T) {
	cmd := newScanClusterCmd()
	if cmd.Use != "cluster [dir]" {
		t.Errorf("Use = %q, want \"cluster [dir]\"", cmd.Use)
	}
	if err := cmd.Args(cmd, []string{"a", "b"}); err == nil {
		t.Error("two args must be rejected")
	}
	if err := cmd.Args(cmd, []string{}); err != nil {
		t.Errorf("zero args rejected: %v", err)
	}
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	if err := cmd.RunE(cmd, []string{writeBlogFixture(t)}); err != nil {
		t.Errorf("RunE on a valid blog dir: %v", err)
	}
	if err := cmd.RunE(cmd, nil); err == nil {
		t.Error("RunE on the default dir (no blog.yaml here) must error")
	}
}
