//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what 글별 큐 후보 조립 — 위반이 있는 글마다 cluster 항목 1건(글당 1건 집약), 무위반 글은 제외
package cluster

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// scanItems builds the queue candidates: one kind=cluster item per article
// with at least one violation. One item per article is the queue's
// idempotent key contract — every violation of the same article shares a key
// by design (Phase010 precedent).
func scanItems(posts []post, b *blogyaml.Blog, lang string, tags, inlinks map[string]int64, edges map[string]bool) []queueio.Item {
	items := make([]queueio.Item, 0)
	for _, p := range posts {
		viols := violations(p, b, tags, inlinks)
		if len(viols) == 0 {
			continue
		}
		items = append(items, clusterItem(p, lang, viols, candidates(p, posts, edges)))
	}
	return items
}
