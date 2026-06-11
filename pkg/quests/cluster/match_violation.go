//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what 재스캔 항목의 violations JSON에서 큐 지정 룰 집합과 겹치는 첫 위반 반환 — 디코드 불가는 미해소로 취급 (cluster-resolved 전용)
package cluster

import (
	"encoding/json"

	scancluster "github.com/park-jun-woo/abloq/pkg/scan/cluster"
)

// matchViolation reports the first violation in the fresh-scan payload whose
// rule the queue item demanded resolved. An unreadable payload (impossible
// for scanner output) counts as unresolved — fail loud, never silently pass.
func matchViolation(raw string, rules map[string]bool) (string, bool) {
	var viols []scancluster.Violation
	if err := json.Unmarshal([]byte(raw), &viols); err != nil {
		return "violations payload unreadable: " + err.Error(), true
	}
	for _, v := range viols {
		if rules[v.Rule] {
			return v.Rule, true
		}
	}
	return "", false
}
