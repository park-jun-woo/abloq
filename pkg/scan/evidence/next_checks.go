//ff:func feature=scan type=rule control=iteration dimension=1 topic=evidence
//ff:what 이번 점검 결과로 상태 갱신 — 실패면 이전 연속 실패 +1, 성공이면 0으로 리셋 (rot 확정의 유일한 입력)
//ff:why 분류(hard/soft)는 표시용 — 카운트는 분류와 무관하게 "실패가 지속되는가"만 본다. 일시 404도 한 번 살아나면 0부터다
package evidence

// nextChecks merges this scan's probe statuses with the previous state into
// the new citation_checks rows. Every currently-cited URL gets a row; rows
// for citations that left the corpus simply stop being updated.
func nextChecks(prev []Check, cites []cite, statuses map[string]string) []Check {
	fails := prevFailures(prev)
	checks := make([]Check, 0, len(cites))
	for _, c := range cites {
		st := statuses[c.URL]
		var n int64
		if st != "ok" {
			n = fails[checkKey(c.URL, c.Lang, c.Section, c.Slug)] + 1
		}
		checks = append(checks, Check{URL: c.URL, Lang: c.Lang, Section: c.Section, Slug: c.Slug,
			Status: st, ConsecutiveFailures: n})
	}
	return checks
}
