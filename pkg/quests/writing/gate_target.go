//ff:func feature=quest type=frame control=sequence
//ff:what Submission의 common.TargetCarrier 구현 — 조립된 단일 글 게이트 Target 반환 (공용 어댑터 계약)
package writing

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// GateTarget exposes the assembled gate target to the shared rule adapter
// (common.TargetCarrier).
func (s *Submission) GateTarget() *agate.Target { return s.Target }
