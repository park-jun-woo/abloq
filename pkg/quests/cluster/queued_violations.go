//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what Seed 시점 고정된 payload violations JSON → 위반 룰 집합 — cluster-resolved의 해소 대상 목록, 부재·공집합은 에러
//ff:why cluster 큐 항목은 위반이 있어야만 발급된다(Phase011) — violations 없는 payload는 변조이거나 발급 버그라 Prepare에서 중단한다
package cluster

import (
	"encoding/json"
	"fmt"

	scancluster "github.com/park-jun-woo/abloq/pkg/scan/cluster"
)

// queuedViolations decodes the frozen queue payload's violations entry (the
// Phase011 scanner's Violation JSON) into the rule set the resolution
// re-check consults. Every cluster queue item carries at least one
// violation, so an empty or absent entry is an error.
func queuedViolations(payload map[string]string) (map[string]bool, error) {
	var viols []scancluster.Violation
	if err := json.Unmarshal([]byte(payload["violations"]), &viols); err != nil {
		return nil, fmt.Errorf("queue payload violations: %w", err)
	}
	if len(viols) == 0 {
		return nil, fmt.Errorf("queue payload violations: empty — a cluster item is only issued for a violating article")
	}
	rules := make(map[string]bool, len(viols))
	for _, v := range viols {
		rules[v.Rule] = true
	}
	return rules, nil
}
