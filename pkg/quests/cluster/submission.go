//ff:type feature=quest type=schema topic=queue
//ff:what 게이트 Context.Submission — 공통 Consumption 임베드 + 대상 글 경로·Key 부품 + payload 위반 룰 집합(해소 재검용)
package cluster

import "github.com/park-jun-woo/abloq/pkg/quests/common"

// Submission is what one cluster-quest submit carries through the gate: the
// shared consumption context (baseline-attached target, working-tree change
// set, queue-scope allowed set extended with the payload candidates), the
// seeded article path, the key parts the re-scan matches on, and the
// violation rules the queue payload demands be resolved. Embedding
// *common.Consumption satisfies the TargetCarrier and ConsCarrier contracts
// by method promotion.
type Submission struct {
	*common.Consumption
	Article   string
	Lang      string
	Section   string
	Slug      string
	ViolRules map[string]bool
}
