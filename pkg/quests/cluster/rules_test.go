//ff:func feature=quest type=frame control=iteration dimension=1 topic=queue
//ff:what Rules 카탈로그 검증 — queue-scope·cluster-resolved 2룰, 전부 LevelFail
package cluster

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRules(t *testing.T) {
	rules := Definition{}.Rules()
	want := []string{"queue-scope", "cluster-resolved"}
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
	}
}
