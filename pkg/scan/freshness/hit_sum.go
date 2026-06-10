//ff:type feature=scan type=schema
//ff:what crawl_hits 합계 1건 — (lang, section, slug)별 봇 히트 총합, CrawlHit.AggSumsJson 행과 1:1
package freshness

// HitSum is one per-article crawl_hits aggregate as the backend's
// CrawlHit.AggSumsJson query emits it. The CLI has no measurement store and
// supplies none (permanent cold start).
type HitSum struct {
	Lang    string `json:"lang"`
	Section string `json:"section"`
	Slug    string `json:"slug"`
	Hits    int64  `json:"hits"`
}
