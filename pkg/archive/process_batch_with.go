//ff:func feature=archive type=client control=sequence
//ff:what 배치 실행 오케스트레이션 — limit 상한 적용 → kind 그룹화 → wayback/indexnow/gsc 제출(사이트 자격 주입) → per-target 결과 합산
//ff:why kind 분기·url 순회·실패 격리는 전부 여기(pkg) 안이다 — SSaC는 무반복·무분기, 에이전트는 외부 API를 직접 치지 않는다 (설계서 §3.3·§4)
package archive

// ProcessBatchWith executes up to limit pending receipts against the three
// external services using the given per-site credentials and returns one
// result Item per processed receipt (done | failed | deferred). Receipts
// beyond limit are not touched — they stay pending for the next run. A
// failed target never aborts the batch.
func ProcessBatchWith(keys Keys, pending []Pending, limit int64) []Item {
	if limit >= 0 && limit < int64(len(pending)) {
		pending = pending[:limit]
	}
	groups := groupByKind(pending)
	results := []Item{}
	results = append(results, processWayback(groups[KindWayback])...)
	results = append(results, processIndexNow(keys.IndexNowKey, groups[KindIndexNow])...)
	results = append(results, processGSC(keys, groups[KindGSCIndex])...)
	return results
}
