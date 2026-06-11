//ff:type feature=quest type=schema topic=queue
//ff:what 게이트 Context.Submission — 공통 Consumption 임베드 + 대상 글 경로 + payload rot URL 목록(해소 재검용)
package evidence

import "github.com/park-jun-woo/abloq/pkg/quests/common"

// Submission is what one evidence-quest submit carries through the gate:
// the shared consumption context (baseline-attached target, working-tree
// change set, queue-scope allowed set, queued claim hashes), the seeded
// article path and the rot URLs the queue payload demands be replaced.
// Embedding *common.Consumption satisfies the TargetCarrier and ConsCarrier
// contracts by method promotion.
type Submission struct {
	*common.Consumption
	Article string
	RotURLs []string
}
