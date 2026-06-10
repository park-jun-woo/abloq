//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what MergeSignals가 기본 Hits 위에 측정 3종을 덮고(search 미반영) 어느 한쪽에만 있는 키도 보존하는지 검증
package report

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

func TestMergeSignals(t *testing.T) {
	base := map[string]priority.Signals{"ko/tech/post-a": {Hits: 9}}
	bots := map[string]Tally{"ko/tech/post-a": {Training: 7, Search: 5, Fetch: 3, MD: 1}}
	pages := map[string]PageTally{"ko/tech/post-b": {Impressions: 40, Clicks: 2}}
	cites := map[string]int64{"ko/tech/post-a": 2}
	m := MergeSignals(base, bots, pages, cites)
	a := m["ko/tech/post-a"]
	if a.Hits != 9 || a.TrainHits != 7 || a.FetcherHits != 3 || a.CitationHits != 2 || a.GSCTrend != 0 {
		t.Errorf("merged signals wrong: %+v", a)
	}
	bRow := m["ko/tech/post-b"]
	if bRow.GSCTrend != 40 || bRow.Hits != 0 {
		t.Errorf("measurement-only key must survive: %+v", bRow)
	}
	if len(m) != 2 {
		t.Errorf("want union of keys, got %d", len(m))
	}
}
