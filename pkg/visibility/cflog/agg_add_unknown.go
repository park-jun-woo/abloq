//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 미지 봇 후보 1건 누적 — 휴리스틱 통과 UA만, 히트 수와 최초/최종 목격 시각 갱신
package cflog

import "github.com/park-jun-woo/abloq/pkg/bots"

// addUnknown accumulates a non-dictionary UA when it passes the bot
// heuristic: hit count plus first/last sighting bounds.
func (a *Agg) addUnknown(rec Record) {
	if !bots.IsBotCandidate(rec.UA) {
		return
	}
	u := a.unknown[rec.UA]
	if u == nil {
		u = &unknownAgg{First: rec.When, Last: rec.When}
		a.unknown[rec.UA] = u
	}
	u.Hits++
	if rec.When.Before(u.First) {
		u.First = rec.When
	}
	if rec.When.After(u.Last) {
		u.Last = rec.When
	}
}
