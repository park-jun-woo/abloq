//ff:func feature=visibility type=generator control=iteration dimension=1 topic=report
//ff:what posts 인덱스 → 글별 표 행 — 측정 집계를 글에 결합, Composite 점수(측정 0이면 date 폴백, Hits 미사용), 우선순위 내림차순 정렬
package report

import (
	"sort"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// rows joins the window aggregates onto the posts index and scores every
// article via rowOf (Composite scorer, date-recency fallback). Order:
// priority descending, then the article key — the report reads as the
// work queue.
func rows(posts []content.Entry, bots map[string]Tally, pages map[string]PageTally, cites map[string]int64, w priority.Weights) []Row {
	scorer := priority.Composite{W: w}
	out := make([]Row, 0, len(posts))
	for _, e := range posts {
		key := queueio.JoinKey(e.Lang, e.Section, e.Slug)
		out = append(out, rowOf(e, bots[key], pages[key], cites[key], scorer))
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Priority != out[j].Priority {
			return out[i].Priority > out[j].Priority
		}
		ki := queueio.JoinKey(out[i].Lang, out[i].Section, out[i].Slug)
		kj := queueio.JoinKey(out[j].Lang, out[j].Section, out[j].Slug)
		return ki < kj
	})
	return out
}
