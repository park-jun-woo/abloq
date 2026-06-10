//ff:func feature=cli type=command control=sequence
//ff:what freshness 명령이 Use/최대 1개 인자를 선언하는지 검증
package main

import "testing"

func TestNewScanFreshnessCmd(t *testing.T) {
	cmd := newScanFreshnessCmd()
	if cmd.Use != "freshness [dir]" {
		t.Errorf("Use = %q, want \"freshness [dir]\"", cmd.Use)
	}
	if err := cmd.Args(cmd, []string{"a", "b"}); err == nil {
		t.Error("two args must be rejected")
	}
	if err := cmd.Args(cmd, []string{}); err != nil {
		t.Errorf("zero args rejected: %v", err)
	}
}
