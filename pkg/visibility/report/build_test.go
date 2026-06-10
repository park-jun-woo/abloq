//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what 리포트 스냅샷 골든 — 고정 Input(명시 ym, now-파생 0)의 markdown·JSON이 testdata 골든과 바이트 동일한지 검증 (UPDATE_GOLDEN=1로 갱신)
package report

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

func TestBuildGolden(t *testing.T) {
	in := Input{
		YM: "2026-04",
		Posts: []content.Entry{
			{Lang: "ko", Section: "tech", Slug: "post-a", Date: "2026-06-01"},
			{Lang: "ko", Section: "tech", Slug: "post-b", Date: "2026-06-03"},
			{Lang: "ko", Section: "tech", Slug: "post-c", Date: "2026-06-05"},
		},
		Bots: []BotSum{
			{Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 7, MDHits: 2},
			{Bot: "ChatGPT-User", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 3},
			{Bot: "PerplexityBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 5},
			{Bot: "ClaudeBot", Lang: "ko", Section: "tech", Slug: "post-b", Hits: 4, MDHits: 1},
		},
		PrevBots: []BotSum{
			{Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 5},
		},
		Pages: []PageSum{
			{Page: "https://fixture.example.com/tech/post-a/", Impressions: 120, Clicks: 8},
			{Page: "https://fixture.example.com/tech/post-b/", Impressions: 40, Clicks: 2},
			{Page: "https://fixture.example.com/", Impressions: 999, Clicks: 99},
		},
		PrevPages: []PageSum{
			{Page: "https://fixture.example.com/tech/post-a/", Impressions: 60, Clicks: 1},
		},
		Cites: []CiteSum{
			{Lang: "ko", Section: "tech", Slug: "post-a", Cited: 2, Total: 3},
		},
		Queue: []QueueCount{
			{Kind: "refresh", Status: "open", Count: 2},
		},
		UnknownBots: []UnknownBot{
			{UA: "PetalBot", Hits: 1},
		},
		URLs: map[string]cflog.Article{
			"/tech/post-a/": {Lang: "ko", Section: "tech", Slug: "post-a"},
			"/tech/post-b/": {Lang: "ko", Section: "tech", Slug: "post-b"},
		},
		Weights: priority.Weights{Fetcher: 3, Train: 1, GSC: 1, Citation: 2},
	}
	r := Build(in)
	// post-a: 3*3 + 1*7 + 1*120 + 2*2 = 140; post-b: 1*4 + 1*40 = 44;
	// post-c: no measurements — cold-start date score (epoch days, a scale
	// of its own by design: no scale conversion on the fallback path).
	if r.Rows[0].Slug != "post-c" || r.Rows[0].Priority != 20609 {
		t.Fatalf("post-c cold-start priority: want 20609 first, got %+v", r.Rows[0])
	}
	if r.Rows[1].Slug != "post-a" || r.Rows[1].Priority != 140 {
		t.Fatalf("post-a priority: want 140, got %+v", r.Rows[1])
	}
	if r.Rows[2].Slug != "post-b" || r.Rows[2].Priority != 44 {
		t.Fatalf("post-b priority: want 44, got %+v", r.Rows[2])
	}
	checkGolden(t, filepath.Join("testdata", "monthly.golden.md"), []byte(Markdown(r)))
	checkGolden(t, filepath.Join("testdata", "monthly.golden.json"), JSON(r))
	// Determinism: a second build is byte-identical.
	if string(JSON(Build(in))) != string(JSON(r)) {
		t.Error("Build must be deterministic")
	}
}
