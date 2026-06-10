//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what rows가 측정 집계를 글에 결합해 가중 점수·우선순위 내림차순(동점은 키 사전순)을 내고 Hits를 점수에 쓰지 않는지 검증
package report

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

func TestRows(t *testing.T) {
	posts := []content.Entry{
		{Lang: "ko", Section: "tech", Slug: "post-a", Date: "2026-06-01"},
		{Lang: "ko", Section: "tech", Slug: "post-b", Date: "2026-06-01"},
		{Lang: "ko", Section: "tech", Slug: "post-c", Date: "2026-06-01"},
	}
	bots := map[string]Tally{
		"ko/tech/post-a": {Training: 7, Search: 5, Fetch: 3, MD: 2},
		"ko/tech/post-b": {Search: 9},
	}
	pages := map[string]PageTally{"ko/tech/post-a": {Impressions: 120, Clicks: 8}}
	cites := map[string]int64{"ko/tech/post-a": 2}
	out := rows(posts, bots, pages, cites, priority.Weights{Fetcher: 3, Train: 1, GSC: 1, Citation: 2})
	// post-b has only search hits — never a scorer input, and Hits stays 0
	// here, so its fallback is the pure date score; post-c ties with it
	// (same date) and the tie breaks on the article key.
	if out[0].Slug != "post-b" || out[1].Slug != "post-c" {
		t.Fatalf("tie order wrong: %s, %s", out[0].Slug, out[1].Slug)
	}
	if out[0].Priority != 20605 || out[1].Priority != 20605 {
		t.Errorf("search-only and empty articles must fall back to the date score: %d, %d", out[0].Priority, out[1].Priority)
	}
	if out[2].Slug != "post-a" || out[2].Priority != 140 {
		t.Errorf("post-a weighted score: want 140, got %+v", out[2])
	}
	if out[2].Search != 5 || out[2].MDHits != 2 || out[2].Clicks != 8 || out[2].Cited != 2 {
		t.Errorf("row display columns wrong: %+v", out[2])
	}
}
