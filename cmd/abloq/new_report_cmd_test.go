//ff:func feature=cli type=command control=sequence topic=report
//ff:what report 부모 명령이 monthly 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewReportCmd(t *testing.T) {
	cmd := newReportCmd()
	if cmd.Use != "report" {
		t.Errorf("Use = %q, want report", cmd.Use)
	}
	sub, _, err := cmd.Find([]string{"monthly"})
	if err != nil || sub.Name() != "monthly" {
		t.Errorf("monthly subcommand missing: %v", err)
	}
}
