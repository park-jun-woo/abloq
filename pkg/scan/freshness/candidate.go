//ff:func feature=scan type=generator control=sequence
//ff:what 신선도 초과 글 1건 → refresh 큐 후보 — 신호 맵에서 측정 신호를 꺼내 date를 채워 스코어러에 전달, payload에 근거(lastmod·threshold)만 (now-파생값 금지)
package freshness

import (
	"strconv"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/queueio"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// candidate builds the refresh queue item for one stale entry. The article's
// signals come from the signals map (the zero value when absent — the
// cold-start path); Date is filled from the entry so the date-recency
// fallback stays available. The payload holds only the generation rationale
// (lastmod + threshold) — never a now-derived value, so re-scans serialize
// byte-identically.
func candidate(e content.Entry, signals map[string]priority.Signals, days int, scorer priority.Scorer) queueio.Item {
	sig := signals[queueio.JoinKey(e.Lang, e.Section, e.Slug)]
	sig.Date = e.Date
	return queueio.Item{
		Kind:     "refresh",
		Slug:     e.Slug,
		Lang:     e.Lang,
		Section:  e.Section,
		Priority: scorer.Score(sig),
		Payload: map[string]string{
			"freshness_days": strconv.Itoa(days),
			"lastmod":        e.Lastmod,
		},
	}
}
