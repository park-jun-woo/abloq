//ff:func feature=scan type=generator control=sequence
//ff:what 신선도 초과 글 1건 → refresh 큐 후보 — 우선순위 산출(콜드스타트: hits→date), payload에 근거(lastmod·threshold)만 (now-파생값 금지)
package freshness

import (
	"strconv"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// candidate builds the refresh queue item for one stale entry. The payload
// holds only the generation rationale (lastmod + threshold) — never a
// now-derived value, so re-scans serialize byte-identically.
func candidate(e content.Entry, hits map[string]int64, days int, scorer priority.Scorer) queueio.Item {
	return queueio.Item{
		Kind:     "refresh",
		Slug:     e.Slug,
		Lang:     e.Lang,
		Section:  e.Section,
		Priority: scorer.Score(priority.Signals{Date: e.Date, Hits: hits[queueio.JoinKey(e.Lang, e.Section, e.Slug)]}),
		Payload: map[string]string{
			"freshness_days": strconv.Itoa(days),
			"lastmod":        e.Lastmod,
		},
	}
}
