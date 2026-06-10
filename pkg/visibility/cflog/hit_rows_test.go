//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what HitRows가 일자→봇→lang→section→slug 사전순으로 행을 고정하는지 검증 (결정적 직렬화)
package cflog

import (
	"reflect"
	"testing"
	"time"
)

func TestHitRows(t *testing.T) {
	urls := map[string]Article{
		"/tech/b/":    {Lang: "ko", Section: "tech", Slug: "b"},
		"/tech/a/":    {Lang: "ko", Section: "tech", Slug: "a"},
		"/en/tech/a/": {Lang: "en", Section: "tech", Slug: "a"},
	}
	agg := NewAgg(urls)
	d1 := time.Date(2026, 6, 1, 1, 0, 0, 0, time.UTC)
	d2 := d1.Add(24 * time.Hour)
	agg.Add(Record{When: d2, URI: "/tech/a/", Status: "200", UA: "GPTBot/1.0"})
	agg.Add(Record{When: d1, URI: "/tech/b/", Status: "200", UA: "GPTBot/1.0"})
	agg.Add(Record{When: d1, URI: "/tech/a/", Status: "200", UA: "ClaudeBot/1.0"})
	agg.Add(Record{When: d1, URI: "/en/tech/a/", Status: "200", UA: "ClaudeBot/1.0"})
	got := agg.HitRows()
	want := []HitRow{
		{HitDate: "2026-06-01", Bot: "ClaudeBot", Lang: "en", Section: "tech", Slug: "a", Hits: 1},
		{HitDate: "2026-06-01", Bot: "ClaudeBot", Lang: "ko", Section: "tech", Slug: "a", Hits: 1},
		{HitDate: "2026-06-01", Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "b", Hits: 1},
		{HitDate: "2026-06-02", Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "a", Hits: 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("HitRows = %+v, want %+v", got, want)
	}
}
