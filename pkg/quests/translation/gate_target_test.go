//ff:func feature=quest type=frame control=sequence
//ff:what GateTarget 검증 — Submission이 조립된 Target을 그대로 내놓아 common.TargetCarrier를 충족하는지
package translation

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

func TestGateTarget(t *testing.T) {
	origin, ko := passPair()
	sub := subWith(t, writeInstance(t), origin, ko)
	var c common.TargetCarrier = sub
	if c.GateTarget() != sub.Target {
		t.Error("GateTarget != Submission.Target")
	}
}
