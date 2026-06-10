//ff:func feature=cli type=command control=sequence
//ff:what scan 부모 명령이 freshness 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewScanCmd(t *testing.T) {
	cmd := newScanCmd()
	if cmd.Use != "scan" {
		t.Errorf("Use = %q, want \"scan\"", cmd.Use)
	}
	sub, _, err := cmd.Find([]string{"freshness"})
	if err != nil || sub.Name() != "freshness" {
		t.Errorf("freshness subcommand missing: %v", err)
	}
}
