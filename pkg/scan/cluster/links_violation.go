//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what [min-internal-links] 확정 아웃링크 수 < geo.min_internal_links 판정 — 코퍼스 실재 간선만 센다
package cluster

import "fmt"

// linksViolation flags an article whose finalized out-degree (resolved,
// in-corpus, deduplicated, no self-link) falls below the threshold.
func linksViolation(outdegree, min int) *Violation {
	if outdegree >= min {
		return nil
	}
	// "below" instead of a less-than sign: encoding/json HTML-escapes that
	// sign as a unicode sequence, making the queue payload unreadable.
	return &Violation{Rule: "min-internal-links", Detail: fmt.Sprintf("outbound internal links %d below min %d", outdegree, min)}
}
