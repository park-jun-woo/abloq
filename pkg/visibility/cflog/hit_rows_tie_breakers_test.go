//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what HitRows의 동률 키(section→slug) 사전순 정렬을 검증
package cflog

import (
	"testing"
	"time"
)

func TestHitRowsTieBreakers(t *testing.T) {
	urls := map[string]Article{
		"/a/x/": {Lang: "ko", Section: "a", Slug: "x"},
		"/b/x/": {Lang: "ko", Section: "b", Slug: "x"},
		"/a/y/": {Lang: "ko", Section: "a", Slug: "y"},
	}
	agg := NewAgg(urls)
	when := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	agg.Add(Record{When: when, URI: "/b/x/", Status: "200", UA: "GPTBot/1.0"})
	agg.Add(Record{When: when, URI: "/a/y/", Status: "200", UA: "GPTBot/1.0"})
	agg.Add(Record{When: when, URI: "/a/x/", Status: "200", UA: "GPTBot/1.0"})
	rows := agg.HitRows()
	order := rows[0].Section + rows[0].Slug + rows[1].Section + rows[1].Slug + rows[2].Section + rows[2].Slug
	if order != "axaybx" {
		t.Errorf("section/slug tie-break order = %q, want axaybx", order)
	}
}
