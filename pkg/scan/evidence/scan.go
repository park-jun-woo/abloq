//ff:func feature=scan type=rule control=sequence topic=evidence
//ff:what 근거 스캔 1회전 — 발행 글 수집 → 인용 일괄 점검 → 연속 실패 갱신 → claims+확정 rot 큐 후보 (CLI·백엔드 공유)
//ff:why prev가 백엔드 상태(citation_checks)의 전부다 — CLI는 prev=nil의 무상태 1회전이라 rot이 절대 확정되지 않고(3회 미만) claims 항목만 남아, 1회 스캔 시점의 endpoint 산출과 바이트 동일하다
package evidence

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Scan runs one evidence pass over the repository at root: unsourced numeric
// claims (gate detector, claims_ignore honored) plus a liveness probe of
// every citation URL. prev is the previous citation state — the backend
// feeds citation_checks, a stateless caller passes nil and gets a
// single-probe report in Result.Checks.
func Scan(root string, b *blogyaml.Blog, prev []Check, ck *Checker) Result {
	arts := articles(root, b)
	cites := collectCitations(arts)
	statuses := ck.CheckAll(citeURLs(cites))
	checks := nextChecks(prev, cites, statuses)
	return Result{Items: scanItems(arts, checks, b.Languages), Checks: checks}
}
