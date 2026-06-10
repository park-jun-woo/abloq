//ff:func feature=visibility type=parser control=iteration dimension=1 topic=report
//ff:what 기본 신호 맵에 측정 신호 3종(분류 크롤·GSC 노출·인용 적중)을 병합 — 어느 쪽에만 있는 키도 보존
package report

import "github.com/park-jun-woo/abloq/pkg/visibility/priority"

// MergeSignals overlays the window measurements onto the base signals map
// (freshness.SignalsMap supplies the all-time Hits). The union of keys is
// kept: an article with measurements but no base entry still gets its
// signals, and vice versa. Search-category hits never enter — Tally keeps
// them for the report table only.
func MergeSignals(base map[string]priority.Signals, bots map[string]Tally, pages map[string]PageTally, cites map[string]int64) map[string]priority.Signals {
	merged := make(map[string]priority.Signals, len(base)+len(bots)+len(pages)+len(cites))
	for k, s := range base {
		merged[k] = s
	}
	for k, t := range bots {
		s := merged[k]
		s.FetcherHits = t.Fetch
		s.TrainHits = t.Training
		merged[k] = s
	}
	for k, p := range pages {
		s := merged[k]
		s.GSCTrend = p.Impressions
		merged[k] = s
	}
	for k, c := range cites {
		s := merged[k]
		s.CitationHits = c
		merged[k] = s
	}
	return merged
}
