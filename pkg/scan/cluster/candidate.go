//ff:type feature=scan type=schema topic=cluster
//ff:what 연결 후보 1건의 큐 payload 표현 — 대상 섹션/slug, 태그 교집합 크기, 제안 방향(out/in)
package cluster

// Candidate is one link suggestion as it rides in the queue payload.
// Directions lists the proposed edges that do not exist yet: "out" (link from
// the violating article to this candidate) and/or "in" (link from this
// candidate back) — an isolated article gets both. SharedTags records the
// ranking rationale so the agent sees why the candidate was suggested.
type Candidate struct {
	Section    string   `json:"section"`
	Slug       string   `json:"slug"`
	SharedTags int64    `json:"shared_tags"`
	Directions []string `json:"directions"`
}
