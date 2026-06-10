//ff:func feature=cli type=command control=iteration dimension=1 topic=report
//ff:what CF 로그 전수 집계 → 현·전월 윈도의 봇별 합계 — hit_date 사전순 윈도 필터, source 비면 빈 집계
package main

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
	"github.com/park-jun-woo/abloq/pkg/visibility/report"
)

// collectLogBotSums aggregates every log object of the source (the
// stateless ingest pass — runIngestCrawl precedent) and splits the hit rows
// into the ym window and the previous month's window by lexicographic
// hit_date comparison. An empty source yields empty sums: the crawl layer
// simply reads zero.
func collectLogBotSums(repo, source, ym string, b *blogyaml.Blog) ([]report.BotSum, []report.BotSum, error) {
	if source == "" {
		return nil, nil, nil
	}
	from, to, err := report.WindowDates(ym)
	if err != nil {
		return nil, nil, err
	}
	// PrevYM of a valid ym is always a valid ym — no second error path.
	prevFrom, prevTo, _ := report.WindowDates(report.PrevYM(ym))
	urls, err := cflog.BuildURLMap(repo, b)
	if err != nil {
		return nil, nil, err
	}
	src, err := cflog.OpenSource(source)
	if err != nil {
		return nil, nil, err
	}
	keys, err := src.List("", "")
	if err != nil {
		return nil, nil, err
	}
	agg, err := cflog.IngestKeys(src, urls, cflog.LogKeys(keys))
	if err != nil {
		return nil, nil, err
	}
	var bots, prevBots []report.BotSum
	for _, r := range agg.HitRows() {
		sum := report.BotSum{Bot: r.Bot, Lang: r.Lang, Section: r.Section, Slug: r.Slug, Hits: r.Hits, MDHits: r.MDHits}
		if r.HitDate >= from && r.HitDate <= to {
			bots = append(bots, sum)
		}
		if r.HitDate >= prevFrom && r.HitDate <= prevTo {
			prevBots = append(prevBots, sum)
		}
	}
	return bots, prevBots, nil
}
