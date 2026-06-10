//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 누적된 히트 카운터를 정준 순서(일자→봇→lang→section→slug 사전순)의 행 목록으로 — 결정적 출력
package cflog

import "sort"

// HitRows flattens the hit aggregation into rows sorted by (hit_date, bot,
// lang, section, slug) so the same logs always serialize identically.
func (a *Agg) HitRows() []HitRow {
	rows := make([]HitRow, 0, len(a.hits))
	for k, c := range a.hits {
		rows = append(rows, HitRow{
			HitDate: k.Date, Bot: k.Bot, Lang: k.Lang, Section: k.Section, Slug: k.Slug,
			Hits: c.Hits, MDHits: c.MDHits,
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		x, y := rows[i], rows[j]
		if x.HitDate != y.HitDate {
			return x.HitDate < y.HitDate
		}
		if x.Bot != y.Bot {
			return x.Bot < y.Bot
		}
		if x.Lang != y.Lang {
			return x.Lang < y.Lang
		}
		if x.Section != y.Section {
			return x.Section < y.Section
		}
		return x.Slug < y.Slug
	})
	return rows
}
