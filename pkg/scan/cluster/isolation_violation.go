//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what [no-isolated-post] 인링크 0 글 판정 — 어떤 글도 가리키지 않는 글은 클러스터 밖이다
package cluster

// isolationViolation flags an article no other article links to.
func isolationViolation(indegree int64) *Violation {
	if indegree > 0 {
		return nil
	}
	return &Violation{Rule: "no-isolated-post", Detail: "no inbound internal links"}
}
