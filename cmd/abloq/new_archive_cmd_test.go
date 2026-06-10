//ff:func feature=cli type=command control=sequence
//ff:what archive 명령이 Use/인자 개수(정확히 1개)를 선언하는지 검증
package main

import "testing"

func TestNewArchiveCmd(t *testing.T) {
	cmd := newArchiveCmd()
	if cmd.Use != "archive <url>" {
		t.Errorf("Use = %q, want \"archive <url>\"", cmd.Use)
	}
	if err := cmd.Args(cmd, []string{}); err == nil {
		t.Error("zero args must be rejected")
	}
	if err := cmd.Args(cmd, []string{"https://a/", "https://b/"}); err == nil {
		t.Error("two args must be rejected")
	}
	if err := cmd.Args(cmd, []string{"https://a/"}); err != nil {
		t.Errorf("one arg rejected: %v", err)
	}
}
