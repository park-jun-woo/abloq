//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what 글 1건과 측정 집계를 표 행 1행으로 결합 — 표시 컬럼 + Composite 점수(측정 0이면 date 폴백, Hits 미사용)
package report

import (
	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// rowOf joins one article's window aggregates into a table row and scores
// it with the Composite scorer. Hits stays 0 here: the report's fallback is
// pure date recency, so search-only traffic can never leak into the score
// through the cold-start path.
func rowOf(e content.Entry, t Tally, p PageTally, c int64, scorer priority.Composite) Row {
	return Row{
		Lang: e.Lang, Section: e.Section, Slug: e.Slug, Date: e.Date,
		Training: t.Training, Search: t.Search, Fetch: t.Fetch, MDHits: t.MD,
		Impressions: p.Impressions, Clicks: p.Clicks, Cited: c,
		Priority: scorer.Score(priority.Signals{
			Date:         e.Date,
			FetcherHits:  t.Fetch,
			TrainHits:    t.Training,
			GSCTrend:     p.Impressions,
			CitationHits: c,
		}),
	}
}
