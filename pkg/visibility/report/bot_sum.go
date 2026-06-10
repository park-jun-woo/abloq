//ff:type feature=visibility type=schema topic=report
//ff:what 윈도 내 봇별·글별 크롤 합계 1행 — CrawlHit.AggBotMonthJson 행과 1:1, 분류는 pkg/bots가 한다 (SQL 분류 금지)
package report

// BotSum is one per-(bot, article) crawl-hit sum over the report window, as
// the backend's CrawlHit.AggBotMonthJson emits it. The bot name stays
// unclassified here: the training/search/fetch split lives only in the
// pkg/bots Go dictionary (BotTotals applies it) — the database has no
// category data, so SQL must never try.
type BotSum struct {
	Bot     string `json:"bot"`
	Lang    string `json:"lang"`
	Section string `json:"section"`
	Slug    string `json:"slug"`
	Hits    int64  `json:"hits"`
	MDHits  int64  `json:"md_hits"`
}
