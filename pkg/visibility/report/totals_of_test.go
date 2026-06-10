//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what TotalsOf가 분류 합·md·노출·클릭·cited를 윈도 전체 합계로 환원하는지 검증
package report

import "testing"

func TestTotalsOf(t *testing.T) {
	got := TotalsOf(
		map[string]Tally{
			"ko/tech/post-a": {Training: 7, Search: 5, Fetch: 3, MD: 2},
			"ko/tech/post-b": {Training: 4, MD: 1},
		},
		map[string]PageTally{
			"ko/tech/post-a": {Impressions: 120, Clicks: 8},
			"ko/tech/post-b": {Impressions: 40, Clicks: 2},
		},
		map[string]int64{"ko/tech/post-a": 2},
	)
	want := Totals{CrawlHits: 19, MDHits: 3, Impressions: 160, Clicks: 10, Cited: 2}
	if got != want {
		t.Errorf("want %+v, got %+v", want, got)
	}
	if (TotalsOf(nil, nil, nil) != Totals{}) {
		t.Error("empty aggregates must total zero")
	}
}
