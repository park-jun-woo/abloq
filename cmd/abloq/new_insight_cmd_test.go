//ff:func feature=cli type=command control=sequence
//ff:what insight 부모 명령이 match 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewInsightCmd(t *testing.T) {
	cmd := newInsightCmd()
	if cmd.Use != "insight" {
		t.Errorf("want Use insight, got %q", cmd.Use)
	}
	if sub, _, err := cmd.Find([]string{"match"}); err != nil || sub.Name() != "match" {
		t.Errorf("want match subcommand registered, got %v / %v", sub, err)
	}
}
