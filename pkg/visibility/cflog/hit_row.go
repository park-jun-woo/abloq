//ff:type feature=visibility type=schema topic=crawl
//ff:what 크롤 히트 집계 1행 — (UTC 일자, 봇, lang, section, slug) 키와 hits/md_hits 누적, crawl_hits 1행 대응
//ff:why JSON 키가 crawl_hits 컬럼명과 일치해야 한다 — 백엔드 @call 래퍼가 이 배열을 그대로 jsonb 업서트에 먹인다. hit_date는 UTC 고정: KST 버킷은 parkjunwoo 개인화고 프레임워크는 블로그 불문 결정적이어야 한다 (Phase012)
package cflog

// HitRow is one aggregated crawl-hit row. HitDate is the UTC log date
// ("2006-01-02"); Hits counts page hits (2xx/304), MDHits counts the
// parallel-served .md hits of the same article key.
type HitRow struct {
	HitDate string `json:"hit_date"`
	Bot     string `json:"bot"`
	Lang    string `json:"lang"`
	Section string `json:"section"`
	Slug    string `json:"slug"`
	Hits    int64  `json:"hits"`
	MDHits  int64  `json:"md_hits"`
}
