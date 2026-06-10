//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what 클러스터 스캔 1회전 — 기본 언어 그래프 구성(노드·실재 간선) → 위반 4종 검출 → 후보 제안 붙인 kind=cluster 큐 후보 (CLI·백엔드 공유)
//ff:why 데이터 소스는 저장소 단일(Phase011 결정) — posts 테이블은 링크 카운트뿐이라 그래프를 만들 수 없고, 두 소스를 섞으면 sync 미선행 시 CLI 등가가 깨진다
package cluster

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// Scan runs one cluster pass over the repository at root: it builds the
// tag/internal-link graph of the default language (translations follow the
// default article's cluster decisions) and returns one kind=cluster queue
// candidate per violating article, with deterministic link suggestions.
func Scan(root string, b *blogyaml.Blog) []queueio.Item {
	if len(b.Languages) == 0 {
		return []queueio.Item{}
	}
	lang := b.Languages[0]
	posts := collectPosts(root, b, lang)
	set := postSet(posts)
	for i := range posts {
		posts[i].Outlinks = filterOutlinks(posts[i], set)
	}
	return scanItems(posts, b, lang, tagCounts(posts), inlinkCounts(posts), edgeSet(posts))
}
