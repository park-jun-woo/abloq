//ff:type feature=visibility type=schema topic=crawl
//ff:what 수집 누적기 — URI 역매핑, (일자,봇,글) 히트 카운터, 미지 봇 누적, 무필터 원시 봇 카운터
package cflog

// Agg accumulates crawl-hit aggregation over parsed records. URLs is the
// repository's URI reverse map; Raw counts every dictionary-bot line with no
// status or mapping filter (the analyze-stats.py comparison point).
type Agg struct {
	URLs    map[string]Article
	hits    map[hitKey]*hitCount
	unknown map[string]*unknownAgg
	Raw     map[string]int64
}
