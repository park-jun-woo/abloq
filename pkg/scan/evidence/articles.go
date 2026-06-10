//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what 스캔 대상 글 수집 — gate.Discover 전 글에서 draft 제외, 발행 코퍼스만 (posts 인덱스와 같은 범위)
package evidence

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/gate"
)

// articles collects the published articles of the repository. The scanner
// watches the live corpus' decay, so drafts — which the posts index also
// skips — are out of scope until they publish.
func articles(root string, b *blogyaml.Blog) []*gate.Article {
	arts := make([]*gate.Article, 0)
	for _, a := range gate.Discover(root, b) {
		if isDraft(a) {
			continue
		}
		arts = append(arts, a)
	}
	return arts
}
