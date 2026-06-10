//ff:func feature=cli type=output control=sequence topic=crawl
//ff:what printCrawlAgg가 히트 행·미지 봇·정렬된 원시 카운터·합계 줄을 결정적으로 출력하는지 검증
package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

func TestPrintCrawlAgg(t *testing.T) {
	urls := map[string]cflog.Article{"/tech/a/": {Lang: "ko", Section: "tech", Slug: "a"}}
	agg := cflog.NewAgg(urls)
	when := time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC)
	agg.Add(cflog.Record{When: when, URI: "/tech/a/", Status: "200", UA: "ClaudeBot/1.0"})
	agg.Add(cflog.Record{When: when, URI: "/tech/a/", Status: "200", UA: "Amazonbot/0.1"})
	agg.Add(cflog.Record{When: when, URI: "/x", Status: "200", UA: "PetalBot/1.0"})
	var out bytes.Buffer
	printCrawlAgg(&out, 2, agg)
	got := out.String()
	want := "crawl hits (date bot lang/section/slug hits md_hits):\n" +
		"  2026-06-01  Amazonbot            ko/tech/a  1 0\n" +
		"  2026-06-01  ClaudeBot            ko/tech/a  1 0\n" +
		"unknown bot candidates (ua hits first_seen last_seen):\n" +
		"  \"PetalBot/1.0\" 1 2026-06-01T12:00:00Z 2026-06-01T12:00:00Z\n" +
		"raw bot counters (no status/mapping filter — analyze-stats.py comparison):\n" +
		"  Amazonbot            1\n" +
		"  ClaudeBot            1\n" +
		"ingest: 2 file(s), 2 article hit(s) (html+md), 1 unknown bot ua(s)\n"
	if got != want {
		t.Errorf("output:\n%q\nwant:\n%q", got, want)
	}
}
