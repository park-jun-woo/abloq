//ff:func feature=cli type=command control=iteration dimension=1
//ff:what image 부모 명령이 og와 convert 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewImageCmd(t *testing.T) {
	cmd := newImageCmd()
	if cmd.Use != "image" {
		t.Errorf("want Use image, got %q", cmd.Use)
	}
	found := map[string]bool{}
	for _, sub := range cmd.Commands() {
		found[sub.Name()] = true
	}
	if !found["og"] || !found["convert"] {
		t.Errorf("want og and convert subcommands, got %v", found)
	}
}
