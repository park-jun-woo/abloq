//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Consumption.GateTarget이 조립된 Target을 그대로 내놓아 TargetCarrier 계약을 충족하는지 검증
package common

import (
	"testing"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func TestConsumptionGateTarget(t *testing.T) {
	tgt := &agate.Target{}
	c := &Consumption{Target: tgt}
	if c.GateTarget() != tgt {
		t.Error("GateTarget must return the assembled target")
	}
	var carrier TargetCarrier = c
	if carrier.GateTarget() != tgt {
		t.Error("Consumption must satisfy TargetCarrier")
	}
}
