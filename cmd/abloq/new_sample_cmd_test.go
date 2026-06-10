//ff:func feature=cli type=command control=sequence topic=citation
//ff:what sample 부모 명령이 citations 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewSampleCmd(t *testing.T) {
	cmd := newSampleCmd()
	if cmd.Use != "sample" {
		t.Errorf("Use = %q, want \"sample\"", cmd.Use)
	}
	sub, _, err := cmd.Find([]string{"citations"})
	if err != nil || sub.Name() != "citations" {
		t.Errorf("citations subcommand missing: %v", err)
	}
}
