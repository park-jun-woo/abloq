//ff:type feature=quest type=schema
//ff:what 테스트 헬퍼 — TargetCarrier 최소 구현체 (Target 1개를 그대로 내놓음)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

type carrier struct{ tgt *agate.Target }
