//ff:func feature=visibility type=parser control=selection topic=report
//ff:what 분류 1건을 집계에 가산 — category별 카운터 분기, md_hits는 분류 불문 누적
package report

// add accumulates one bot sum into the tally under its dictionary category.
func (t *Tally) add(category string, hits, mdHits int64) {
	t.MD += mdHits
	switch category {
	case "training":
		t.Training += hits
	case "search":
		t.Search += hits
	case "fetch":
		t.Fetch += hits
	}
}
