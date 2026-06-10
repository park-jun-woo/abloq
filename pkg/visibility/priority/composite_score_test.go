//ff:func feature=visibility type=scorer control=iteration dimension=1
//ff:what Composite.Score가 측정 전부 0이면 ColdStart 값을 무변형 반환(date 20607/20605 보존)하고 측정이 있으면 가중 합인지 검증
package priority

import "testing"

func TestCompositeScore(t *testing.T) {
	c := Composite{W: Weights{Fetcher: 3, Train: 1, GSC: 1, Citation: 2}}
	cases := []struct {
		name string
		s    Signals
		want int64
	}{
		// Cold-start fallback: the date epoch-day scores pass through
		// untouched — the scenario-freshness asserts pin these exact values.
		{"cold start date", Signals{Date: "2026-06-03"}, 20607},
		{"cold start older date", Signals{Date: "2026-06-01"}, 20605},
		{"cold start hits", Signals{Date: "2026-06-01", Hits: 42}, 42},
		// Any non-zero measurement signal routes to the measured score.
		{"measured wins", Signals{Date: "2026-06-01", Hits: 42, FetcherHits: 3, TrainHits: 7, GSCTrend: 120, CitationHits: 2}, 140},
		{"single measured signal", Signals{Date: "2026-06-01", Hits: 42, GSCTrend: 40}, 40},
	}
	for _, tc := range cases {
		if got := c.Score(tc.s); got != tc.want {
			t.Errorf("%s: want %d, got %d", tc.name, tc.want, got)
		}
	}
}
