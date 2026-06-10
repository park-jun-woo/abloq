//ff:type feature=scan type=schema topic=cluster
//ff:what 정렬 중인 후보 1건 — payload 후보에 정렬 신호(동일 섹션·발행일 거리·동률 키)를 붙인 내부 표현
package cluster

// ranked carries one candidate through the sort: the payload form plus the
// ranking signals that never reach the queue file.
type ranked struct {
	cand        Candidate
	sameSection bool
	dateDist    int64
	key         string
}
