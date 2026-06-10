//ff:type feature=visibility type=schema topic=report
//ff:what 글 1건의 분류별 크롤 집계 — training/search/fetch 히트와 md_hits (분류는 pkg/bots 사전)
package report

// Tally is one article's classified crawl aggregate. Search rides in the
// report table only — it never reaches a Scorer (§6.1 names train bots and
// fetchers as the priority signals). MD accumulates the .md twin hits over
// every bot category.
type Tally struct {
	Training int64
	Search   int64
	Fetch    int64
	MD       int64
}
