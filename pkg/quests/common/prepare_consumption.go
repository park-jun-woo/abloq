//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Prepare 공통 본체 — 기준선 Target 조립 + porcelain 변경 집합 + 허용 경로 집합으로 Consumption 구성, 글 원문 바이트 반환 (소비 퀘스트 3종 공유)
package common

// PrepareConsumption assembles the shared consumption context for one
// submission: the baseline-attached single-article target, the working-tree
// change set captured now, and the queue-scope allowed set expanded from the
// payload's per-language keys. Every error here aborts the submit without
// burning a try (instance state, not authoring quality).
func PrepareConsumption(p QueuePayload) (*Consumption, []byte, error) {
	tgt, body, err := AssembleBaseTarget(p.Root, p.Article, p.Lang, p.Section, p.Slug)
	if err != nil {
		return nil, nil, err
	}
	changed, err := ChangedSet(p.Root)
	if err != nil {
		return nil, nil, err
	}
	return &Consumption{Target: tgt, Changed: changed, Allowed: AllowedPaths(p.Keys)}, body, nil
}
