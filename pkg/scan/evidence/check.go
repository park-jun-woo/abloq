//ff:type feature=scan type=schema topic=evidence
//ff:what 인용 URL 점검 상태 1건 — (url, lang, section, slug) 키 + 판정(ok/hard/soft)·연속 실패 수, citation_checks 1행과 1:1
//ff:why rot 확정은 분류와 무관하게 연속 실패 3회 — 일시 404·CDN 오류 오탐을 막는 백엔드 상태이며, JSON 키가 jsonb_agg 공급과 일치해야 한다
package evidence

// Check is the link-rot state of one citation occurrence: the same URL cited
// by two articles is two rows, because the fix (replace or drop the citation)
// happens per article. Status classifies the last probe — "ok", "hard"
// (404/410/dead domain) or "soft" (5xx/timeout/network) — but rot confirmation
// reads ConsecutiveFailures only. JSON keys mirror the citation_checks columns.
type Check struct {
	URL                 string `json:"url"`
	Lang                string `json:"lang"`
	Section             string `json:"section"`
	Slug                string `json:"slug"`
	Status              string `json:"status"`
	ConsecutiveFailures int64  `json:"consecutive_failures"`
}
