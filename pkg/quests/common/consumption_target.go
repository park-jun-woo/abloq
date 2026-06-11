//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Consumption의 common.TargetCarrier 구현 — 기준선 부착 게이트 Target 반환, 임베드한 Submission에 승격 (공용 어댑터 계약)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// GateTarget exposes the baseline-attached gate target to the shared rule
// adapter (TargetCarrier); quest submissions inherit it by embedding.
func (c *Consumption) GateTarget() *agate.Target { return c.Target }
