//ff:type feature=visibility type=schema topic=crawl
//ff:what 히트 키 1개의 카운터 — 페이지 히트와 .md 병행 서빙 히트 분리 누적
package cflog

// hitCount accumulates one key's page and .md hits.
type hitCount struct {
	Hits   int64
	MDHits int64
}
