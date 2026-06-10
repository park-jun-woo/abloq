//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 히트 행 목록의 hits+md_hits 총합 — 수집 응답의 hits 스칼라
package cflog

// TotalHits sums Hits+MDHits over all rows: .md consumption is direct agent
// traffic, so the ingest response counts it alongside page hits.
func TotalHits(rows []HitRow) int64 {
	var total int64
	for _, r := range rows {
		total += r.Hits + r.MDHits
	}
	return total
}
