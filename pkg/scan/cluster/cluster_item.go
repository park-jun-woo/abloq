//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what 글 1편 → kind=cluster 큐 항목 — payload에 violations(JSON)·candidates(JSON), priority는 위반+후보 건수 합
//ff:why payload 값은 문자열만(JSONB 바이트 왕복) — 구조 데이터는 JSON 문자열로 한 번 더 감싸 CLI/endpoint 직렬화가 갈라질 수 없게 한다 (evidence 선례)
package cluster

import (
	"encoding/json"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// clusterItem builds the queue candidate for one violating article.
// Violations and link candidates ride as JSON strings inside the string-only
// payload (struct field order keeps the bytes deterministic); priority is
// the finding-plus-suggestion count — the more curation work, the sooner the
// agent should look. No now-derived value enters the payload.
func clusterItem(p post, lang string, viols []Violation, cands []Candidate) queueio.Item {
	payload := map[string]string{}
	data, _ := json.Marshal(viols) // cannot fail: plain string fields
	payload["violations"] = string(data)
	if len(cands) > 0 {
		data, _ := json.Marshal(cands) // cannot fail: strings and int64 only
		payload["candidates"] = string(data)
	}
	return queueio.Item{
		Kind:     "cluster",
		Slug:     p.Slug,
		Lang:     lang,
		Section:  p.Section,
		Priority: int64(len(viols) + len(cands)),
		Payload:  payload,
	}
}
