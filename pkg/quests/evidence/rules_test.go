//ff:func feature=quest type=frame control=iteration dimension=1 topic=queue
//ff:what Rules 카탈로그 검증 — 7룰의 ID·순서·전부 LevelFail, lastmod 강제 룰 비편입(임계 미달 변경의 lastmod 불변경 정책)
package evidence

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRules(t *testing.T) {
	rules := Definition{}.Rules()
	want := []string{
		"claims-resolved", "rot-resolved", "claim-scope", "queue-scope",
		"numeric-claim-sourced", "min-sources", "citation-exists",
	}
	if len(rules) != len(want) {
		t.Fatalf("rule count = %d, want %d", len(rules), len(want))
	}
	for i, r := range rules {
		if r.Meta.ID != want[i] {
			t.Errorf("rules[%d] = %s, want %s", i, r.Meta.ID, want[i])
		}
		if r.Meta.Level != rgate.LevelFail {
			t.Errorf("%s: level must be Fail (Review deadlocks the ratchet)", r.Meta.ID)
		}
		if r.Meta.ID == "honest-lastmod" || r.Meta.ID == "lastmod-advance" {
			t.Error("evidence must not force a lastmod update (sub-threshold honesty)")
		}
	}
}
