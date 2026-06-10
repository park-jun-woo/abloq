//ff:func feature=queueio type=generator control=iteration dimension=1
//ff:what 행 목록 → DB id 목록 (상태 전이 쿼리 입력, nil 대신 빈 슬라이스 — JSON "[]" 보장)
package queueio

// rowIDs collects the database ids of rows. The result is never nil so that
// json.Marshal yields "[]" — the transition queries expect a JSON array.
func rowIDs(rows []Row) []int64 {
	ids := make([]int64, 0, len(rows))
	for _, r := range rows {
		ids = append(ids, r.ID)
	}
	return ids
}
