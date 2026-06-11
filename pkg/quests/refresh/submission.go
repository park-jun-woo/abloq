//ff:type feature=quest type=schema topic=queue
//ff:what 게이트 Context.Submission — 공통 Consumption(기준선 Target·변경 집합·허용 집합) 임베드 + 대상 글 경로
package refresh

import "github.com/park-jun-woo/abloq/pkg/quests/common"

// Submission is what one refresh-quest submit carries through the gate: the
// shared consumption context (baseline-attached target, working-tree change
// set, queue-scope allowed set) plus the seeded article path. Embedding
// *common.Consumption satisfies both the TargetCarrier and ConsCarrier
// contracts by method promotion.
type Submission struct {
	*common.Consumption
	Article string
}
