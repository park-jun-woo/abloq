//ff:func feature=cli type=output control=iteration dimension=1 topic=crawl
//ff:what crawl 수집 결과 출력 — 히트 행(일자 봇 글 hits/md), 미지 봇 후보, 무필터 원시 봇 카운터(대조 지점), 합계 한 줄
package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

// printCrawlAgg prints one stateless crawl aggregation: the per-article hit
// rows, the unknown-bot candidates and the raw per-bot counters with no
// status or mapping filter — the analyze-stats.py comparison point.
func printCrawlAgg(out io.Writer, files int, agg *cflog.Agg) {
	hits := agg.HitRows()
	fmt.Fprintln(out, "crawl hits (date bot lang/section/slug hits md_hits):")
	for _, r := range hits {
		fmt.Fprintf(out, "  %s  %-20s %s/%s/%s  %d %d\n", r.HitDate, r.Bot, r.Lang, r.Section, r.Slug, r.Hits, r.MDHits)
	}
	unknown := agg.UnknownRows()
	fmt.Fprintln(out, "unknown bot candidates (ua hits first_seen last_seen):")
	for _, u := range unknown {
		fmt.Fprintf(out, "  %q %d %s %s\n", u.UA, u.Hits, u.FirstSeen, u.LastSeen)
	}
	fmt.Fprintln(out, "raw bot counters (no status/mapping filter — analyze-stats.py comparison):")
	names := make([]string, 0, len(agg.Raw))
	for name := range agg.Raw {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(out, "  %-20s %d\n", name, agg.Raw[name])
	}
	fmt.Fprintf(out, "ingest: %d file(s), %d article hit(s) (html+md), %d unknown bot ua(s)\n",
		files, cflog.TotalHits(hits), len(unknown))
}
