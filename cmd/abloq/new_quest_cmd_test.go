//ff:func feature=cli type=command control=iteration dimension=1
//ff:what newQuestCmd 검증 — writing 퀘스트가 마운트되고 reins 표준 서브커맨드(scan/next/submit/status/export/rules)가 달리는지
package main

import "testing"

func TestNewQuestCmd(t *testing.T) {
	cmd := newQuestCmd()
	if cmd.Use != "quest" {
		t.Fatalf("Use = %q, want quest", cmd.Use)
	}
	sub, _, err := cmd.Find([]string{"writing"})
	if err != nil || sub.Name() != "writing" {
		t.Fatalf("writing subcommand missing: %v", err)
	}
	for _, name := range []string{"scan", "next", "submit", "status", "export", "rules"} {
		c, _, err := cmd.Find([]string{"writing", name})
		if err != nil || c.Name() != name {
			t.Errorf("writing %s: missing (%v)", name, err)
		}
	}
}
