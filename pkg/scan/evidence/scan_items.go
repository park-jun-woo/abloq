//ff:func feature=scan type=generator control=iteration dimension=1 topic=evidence
//ff:what 글별 큐 후보 조립 — 무출처 주장 또는 확정 rot이 있는 글마다 evidence 항목 1건, 둘 다 없으면 제외
package evidence

import (
	"github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// scanItems builds the queue candidates: one kind=evidence item per article
// that has unsourced claims, confirmed rot, or both. One item per article is
// the queue's idempotent key contract — claims and rot of the same article
// share a key by design (final consistency over payload merge).
func scanItems(arts []*gate.Article, checks []Check, langs []string) []queueio.Item {
	items := make([]queueio.Item, 0)
	for _, a := range arts {
		claims := unsourcedClaims(a)
		rots := rotURLs(checks, a.Lang, a.Section, a.Slug)
		if len(claims) == 0 && len(rots) == 0 {
			continue
		}
		items = append(items, evidenceItem(a, claims, rots, langs))
	}
	return items
}
