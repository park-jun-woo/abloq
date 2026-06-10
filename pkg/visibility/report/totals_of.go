//ff:func feature=visibility type=parser control=iteration dimension=1 topic=report
//ff:what 글별 집계 맵 3종을 윈도 전체 합계로 환원 — 분류 합(training+search+fetch)·md·노출·클릭·cited
package report

// TotalsOf folds the per-article aggregates into the window total. Crawl
// hits sum every category (the table splits them; the total is the raw
// volume); GSC counts only article-attributed pages, mirroring the rows.
func TotalsOf(bots map[string]Tally, pages map[string]PageTally, cites map[string]int64) Totals {
	var t Totals
	for _, b := range bots {
		t.CrawlHits += b.Training + b.Search + b.Fetch
		t.MDHits += b.MD
	}
	for _, p := range pages {
		t.Impressions += p.Impressions
		t.Clicks += p.Clicks
	}
	for _, c := range cites {
		t.Cited += c
	}
	return t
}
