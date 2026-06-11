//ff:func feature=archive type=client control=sequence
//ff:what 빈 Keys로 ProcessBatchWith에 위임 — 전역 env 자격을 쓰는 단일 사이트 하위호환 진입점
package archive

// ProcessBatch keeps the single-site entrypoint: empty Keys make every
// client fall back to the instance-global environment credentials
// (INDEXNOW_KEY, GSC_SA_JSON / GSC_SA_JSON_PATH).
func ProcessBatch(pending []Pending, limit int64) []Item {
	return ProcessBatchWith(Keys{}, pending, limit)
}
