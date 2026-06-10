//ff:func feature=archive type=client control=sequence
//ff:what GSC 그룹 처리 — 신규(date==lastmod) 우선 정렬 → 쿼터 상한까지 제출·초과분 deferred → SA 토큰 실패 시 전 항목 failed
//ff:why GSC Indexing API 쿼터(일 200) 보호: 분할·우선순위 판단은 전부 pkg 내부 — 초과분은 deferred로 이월하고 retry/다음 process 주기가 흡수한다 (Phase008)
package archive

import "encoding/json"

// processGSC submits the targets to the Indexing API, new posts first.
// Submissions beyond the per-run quota are deferred (not failed) — a later
// retry + process cycle absorbs them without burning quota now.
func processGSC(pending []Pending) []Item {
	if len(pending) == 0 {
		return nil
	}
	endpoint := envOr("GSC_API_BASE", "https://indexing.googleapis.com") + "/v3/urlNotifications:publish"
	ordered := gscOrder(pending)
	head, tail := ordered, []Pending(nil)
	if quota := gscQuota(); len(ordered) > quota {
		head, tail = ordered[:quota], ordered[quota:]
	}
	deferred := fanoutItems(tail, endpoint,
		json.RawMessage(`{"reason":"gsc daily quota split — deferred to a later run"}`), StatusDeferred)
	token, err := GSCToken(ScopeIndexing)
	if err != nil {
		return append(fanoutItems(head, endpoint, wrapResponse(0, []byte(err.Error())), StatusFailed), deferred...)
	}
	return append(gscSubmitAll(head, endpoint, token), deferred...)
}
