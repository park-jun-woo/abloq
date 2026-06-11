//ff:func feature=quest type=frame control=sequence
//ff:what 테스트 헬퍼 — carrier의 TargetCarrier 구현(GateTarget)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

func (c *carrier) GateTarget() *agate.Target { return c.tgt }
