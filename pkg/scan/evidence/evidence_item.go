//ff:func feature=scan type=generator control=sequence topic=evidence
//ff:what 글 1편 → kind=evidence 큐 항목 — payload에 claims(JSON)·rot_urls(JSON), priority는 검출 건수 합
//ff:why payload 값은 문자열만(JSONB 바이트 왕복) — 구조 데이터는 JSON 문자열로 한 번 더 감싸 CLI/endpoint 직렬화가 갈라질 수 없게 한다
package evidence

import (
	"encoding/json"

	"github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// evidenceItem builds the queue candidate for one article. Claims and rot
// URLs ride as JSON strings inside the string-only payload (struct field
// order keeps the bytes deterministic); priority is the finding count — the
// more decay, the sooner the agent should look.
func evidenceItem(a *gate.Article, claims []ClaimRef, rots []string) queueio.Item {
	payload := map[string]string{}
	if len(claims) > 0 {
		data, _ := json.Marshal(claims) // cannot fail: plain string fields
		payload["claims"] = string(data)
	}
	if len(rots) > 0 {
		data, _ := json.Marshal(rots) // cannot fail: []string
		payload["rot_urls"] = string(data)
	}
	return queueio.Item{
		Kind:     "evidence",
		Slug:     a.Slug,
		Lang:     a.Lang,
		Section:  a.Section,
		Priority: int64(len(claims) + len(rots)),
		Payload:  payload,
	}
}
