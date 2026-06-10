//ff:func feature=scan type=rule control=sequence
//ff:what Scan이 freshness_days 초과 글만 refresh 후보로 만들고 payload(섹션 제외 근거)·콜드스타트 우선순위를 채우는지 검증
package freshness

import (
	"testing"
	"time"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

func TestScan(t *testing.T) {
	now := time.Date(2026, 6, 11, 0, 0, 0, 0, time.UTC)
	entries := []content.Entry{
		{Lang: "ko", Section: "tech", Slug: "stale-a", Date: "2026-06-01", Lastmod: "2026-06-05"},
		{Lang: "ko", Section: "tech", Slug: "fresh", Date: "2026-06-10", Lastmod: "2026-06-10"},
		{Lang: "ko", Section: "tech", Slug: "bad-date", Date: "", Lastmod: "unparseable"},
	}
	items := Scan(entries, map[string]int64{}, 1, now, priority.ColdStart{})
	if len(items) != 1 {
		t.Fatalf("want 1 stale item, got %d: %+v", len(items), items)
	}
	it := items[0]
	if it.Kind != "refresh" || it.Slug != "stale-a" || it.Section != "tech" {
		t.Errorf("unexpected item: %+v", it)
	}
	if it.Payload["lastmod"] != "2026-06-05" || it.Payload["freshness_days"] != "1" {
		t.Errorf("payload must carry the rationale: %+v", it.Payload)
	}
	if it.Priority != 20605 {
		t.Errorf("cold-start priority must be the date score: %d", it.Priority)
	}
	// Crawl hits override the date fallback.
	hits := map[string]int64{"ko/tech/stale-a": 42}
	items = Scan(entries, hits, 1, now, priority.ColdStart{})
	if items[0].Priority != 42 {
		t.Errorf("hits sum must win: %d", items[0].Priority)
	}
}
