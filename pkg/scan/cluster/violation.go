//ff:type feature=scan type=schema topic=cluster
//ff:what 클러스터 위반 1건의 큐 payload 표현 — 룰ID(4종)와 사람이 읽는 근거 한 줄
package cluster

// Violation is one cluster finding as it rides in the queue payload. Rule is
// one of tag-taxonomy, no-orphan-tag, min-internal-links, no-isolated-post;
// Detail is the human-readable rationale (no now-derived values — the queue
// file serialization must stay deterministic).
type Violation struct {
	Rule   string `json:"rule"`
	Detail string `json:"detail"`
}
