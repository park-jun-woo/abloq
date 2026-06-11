//ff:func feature=cli type=command control=iteration dimension=1
//ff:what newQuestCmd 검증 — 퀘스트 5종이 마운트되고 reins 표준 서브커맨드(scan/next/submit/status/export/rules)가 달리는지
package main

import "testing"

func TestNewQuestCmd(t *testing.T) {
	cmd := newQuestCmd()
	if cmd.Use != "quest" {
		t.Fatalf("Use = %q, want quest", cmd.Use)
	}
	for _, quest := range []string{"writing", "translation", "refresh", "evidence", "cluster"} {
		sub, _, err := cmd.Find([]string{quest})
		if err != nil || sub.Name() != quest {
			t.Fatalf("%s subcommand missing: %v", quest, err)
		}
	}
	for _, name := range []string{"scan", "next", "submit", "status", "export", "rules"} {
		c, _, err := cmd.Find([]string{"refresh", name})
		if err != nil || c.Name() != name {
			t.Errorf("refresh %s: missing (%v)", name, err)
		}
	}
}
