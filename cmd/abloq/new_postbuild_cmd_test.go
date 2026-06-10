//ff:func feature=cli type=command control=iteration dimension=1
//ff:what postbuild 부모 명령이 md 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewPostbuildCmd(t *testing.T) {
	cmd := newPostbuildCmd()
	if cmd.Use != "postbuild" {
		t.Errorf("want Use postbuild, got %q", cmd.Use)
	}
	found := false
	for _, sub := range cmd.Commands() {
		if sub.Name() == "md" {
			found = true
		}
	}
	if !found {
		t.Errorf("want md subcommand registered, got %v", cmd.Commands())
	}
}
