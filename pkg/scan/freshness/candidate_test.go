//ff:func feature=scan type=generator control=sequence
//ff:what candidate가 refresh kind·근거 payload·우선순위(신호 맵의 hits 우선, date 폴백)를 채우는지 검증
package freshness

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

func TestCandidate(t *testing.T) {
	e := content.Entry{Lang: "ko", Section: "tech", Slug: "post-a", Date: "2026-06-01", Lastmod: "2026-06-05"}
	signals := map[string]priority.Signals{"ko/tech/post-a": {Hits: 3}}
	it := candidate(e, signals, 90, priority.ColdStart{})
	if it.Kind != "refresh" || it.Priority != 3 {
		t.Errorf("unexpected item: %+v", it)
	}
	if it.Payload["freshness_days"] != "90" || it.Payload["lastmod"] != "2026-06-05" {
		t.Errorf("rationale payload wrong: %+v", it.Payload)
	}
	if it.Payload["section"] != "" {
		t.Error("section must stay a first-class field, not payload")
	}
	// Measured signals reach the scorer through the same map.
	signals["ko/tech/post-a"] = priority.Signals{FetcherHits: 2, GSCTrend: 10}
	it = candidate(e, signals, 90, priority.Composite{W: priority.Weights{Fetcher: 3, GSC: 1}})
	if it.Priority != 16 {
		t.Errorf("measured priority must be the weighted sum: %d", it.Priority)
	}
}
