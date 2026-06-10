//ff:func feature=visibility type=scorer control=iteration dimension=1
//ff:what Measured.Score의 가중 합 골든 값 — 신호별 계수 반영, search 신호 부재, 0 가중치 무시를 검증
package priority

import "testing"

func TestMeasuredScore(t *testing.T) {
	w := Weights{Fetcher: 3, Train: 1, GSC: 1, Citation: 2}
	cases := []struct {
		name string
		s    Signals
		want int64
	}{
		{"all signals", Signals{FetcherHits: 3, TrainHits: 7, GSCTrend: 120, CitationHits: 2}, 140},
		{"fetcher only", Signals{FetcherHits: 5}, 15},
		{"train only", Signals{TrainHits: 4}, 4},
		{"gsc only", Signals{GSCTrend: 40}, 40},
		{"citation only", Signals{CitationHits: 3}, 6},
		{"zero signals", Signals{Date: "2026-06-01", Hits: 9}, 0},
	}
	for _, tc := range cases {
		if got := (Measured{W: w}).Score(tc.s); got != tc.want {
			t.Errorf("%s: want %d, got %d", tc.name, tc.want, got)
		}
	}
	// A zero weight silences its signal.
	got := (Measured{W: Weights{Fetcher: 3}}).Score(Signals{FetcherHits: 2, GSCTrend: 100})
	if got != 6 {
		t.Errorf("zero gsc weight must drop the gsc signal: %d", got)
	}
}
