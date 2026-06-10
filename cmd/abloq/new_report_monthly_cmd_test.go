//ff:func feature=cli type=command control=sequence topic=report
//ff:what report monthly 명령이 --ym/--source 플래그와 최대 1개 인자를 선언하는지 검증
package main

import (
	"testing"
)

func TestNewReportMonthlyCmd(t *testing.T) {
	cmd := newReportMonthlyCmd()
	if cmd.Flags().Lookup("ym") == nil || cmd.Flags().Lookup("source") == nil {
		t.Error("--ym and --source flags must exist")
	}
	if err := cmd.Args(cmd, []string{"a", "b"}); err == nil {
		t.Error("two args must be rejected")
	}
	if err := cmd.Args(cmd, []string{}); err != nil {
		t.Errorf("zero args rejected: %v", err)
	}
}
