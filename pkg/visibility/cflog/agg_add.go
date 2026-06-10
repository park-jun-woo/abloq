//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 레코드 1건 누적 — 사전 봇이면 원시 카운터 +1 후 2xx/304·글 매핑 통과분만 hits/md_hits 가산, 사전 밖 봇 후보는 미지 봇 누적
package cflog

import "github.com/park-jun-woo/abloq/pkg/bots"

// Add accumulates one parsed record: dictionary bots always bump the raw
// counter; only 2xx/304 hits that map to an article land in the hit
// aggregation (md_hits for the parallel-served .md path, hits otherwise).
// Non-dictionary UAs that look like bots feed the unknown-bot candidates.
func (a *Agg) Add(rec Record) {
	name, known := bots.Classify(rec.UA)
	if !known {
		a.addUnknown(rec)
		return
	}
	a.Raw[name]++
	if !statusOK(rec.Status) {
		return
	}
	art, mapped := a.URLs[rec.URI]
	if !mapped {
		return
	}
	key := hitKey{Date: rec.When.Format("2006-01-02"), Bot: name, Lang: art.Lang, Section: art.Section, Slug: art.Slug}
	c := a.hits[key]
	if c == nil {
		c = &hitCount{}
		a.hits[key] = c
	}
	if art.MD {
		c.MDHits++
	} else {
		c.Hits++
	}
}
