//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what Add가 사전 봇만 원시 카운터에, 2xx/304·매핑 통과분만 hits/md_hits에 누적하는지 검증 — 404·미매핑·사람 UA 제외
package cflog

import (
	"testing"
	"time"
)

func TestAggAdd(t *testing.T) {
	urls := map[string]Article{
		"/tech/a/":   {Lang: "ko", Section: "tech", Slug: "a"},
		"/tech/a.md": {Lang: "ko", Section: "tech", Slug: "a", MD: true},
	}
	when := time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC)
	gpt := "Mozilla/5.0; compatible; GPTBot/1.2"
	agg := NewAgg(urls)
	agg.Add(Record{When: when, URI: "/tech/a/", Status: "200", UA: gpt})
	agg.Add(Record{When: when, URI: "/tech/a.md", Status: "200", UA: gpt})
	agg.Add(Record{When: when, URI: "/tech/a/", Status: "404", UA: gpt})
	agg.Add(Record{When: when, URI: "/style.css", Status: "200", UA: gpt})
	agg.Add(Record{When: when, URI: "/tech/a/", Status: "200", UA: "Mozilla/5.0 Chrome/125 Safari/537.36"})
	if agg.Raw["GPTBot"] != 4 {
		t.Errorf("Raw[GPTBot] = %d, want 4 (no status/mapping filter)", agg.Raw["GPTBot"])
	}
	rows := agg.HitRows()
	if len(rows) != 1 || rows[0].Hits != 1 || rows[0].MDHits != 1 {
		t.Errorf("rows = %+v, want one a-row with hits=1 md_hits=1", rows)
	}
	if len(agg.UnknownRows()) != 0 {
		t.Errorf("browser UA landed in unknown bots")
	}
}
