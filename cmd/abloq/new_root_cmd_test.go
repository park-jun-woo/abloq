//ff:func feature=cli type=command control=iteration dimension=1
//ff:what abloq 루트 명령이 abloq라는 이름으로 생성되고 validate 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewRootCmd(t *testing.T) {
	cmd := newRootCmd()
	if cmd.Use != "abloq" {
		t.Errorf("want Use abloq, got %q", cmd.Use)
	}
	if !cmd.SilenceUsage || !cmd.SilenceErrors {
		t.Errorf("want SilenceUsage and SilenceErrors true, got %v / %v", cmd.SilenceUsage, cmd.SilenceErrors)
	}
	found := false
	for _, sub := range cmd.Commands() {
		if sub.Name() == "validate" {
			found = true
		}
	}
	if !found {
		t.Errorf("want validate subcommand registered, got %v", cmd.Commands())
	}
}
