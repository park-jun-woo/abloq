//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what rowOf가 글 1건의 표시 컬럼을 결합하고 측정 가중 점수(Hits 미사용)·측정 0이면 date 폴백 점수를 내는지 검증
package report

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

func TestRowOf(t *testing.T) {
	scorer := priority.Composite{W: priority.Weights{Fetcher: 3, Train: 1, GSC: 1, Citation: 2}}
	e := content.Entry{Lang: "ko", Section: "tech", Slug: "post-a", Date: "2026-06-01"}
	r := rowOf(e, Tally{Training: 7, Search: 5, Fetch: 3, MD: 2}, PageTally{Impressions: 120, Clicks: 8}, 2, scorer)
	if r.Lang != "ko" || r.Section != "tech" || r.Slug != "post-a" || r.Date != "2026-06-01" {
		t.Errorf("identity columns wrong: %+v", r)
	}
	if r.Training != 7 || r.Search != 5 || r.Fetch != 3 || r.MDHits != 2 || r.Impressions != 120 || r.Clicks != 8 || r.Cited != 2 {
		t.Errorf("measurement columns wrong: %+v", r)
	}
	// 3*3 + 1*7 + 1*120 + 2*2 = 140 — search hits never enter the score.
	if r.Priority != 140 {
		t.Errorf("weighted priority: want 140, got %d", r.Priority)
	}
	// No measurements — the cold-start date score (epoch days).
	empty := rowOf(e, Tally{}, PageTally{}, 0, scorer)
	if empty.Priority != 20605 {
		t.Errorf("cold-start fallback: want 20605, got %d", empty.Priority)
	}
}
