//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what Seed — 원문(기본 언어) 글 경로 인자마다 (선언 언어 − 기본 언어) × Item 시드, lastmod 비교로 최신 번역은 미생성, 인자 없음은 에러
package translation

import (
	"fmt"

	"github.com/park-jun-woo/reins/pkg/quest"
)

// Seed creates the language-matrix TODO items: one item per (origin article
// argument × declared non-default language), skipping translations whose
// lastmod is already >= the origin's (only stale or missing ones enter the
// ratchet).
func (Definition) Seed(args []string) ([]*quest.Item, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("usage: scan <default-language article.md> [article.md...]")
	}
	var items []*quest.Item
	for _, arg := range args {
		src, err := seedOrigin(arg)
		if err != nil {
			return nil, err
		}
		its, err := seedItems(src)
		if err != nil {
			return nil, err
		}
		items = append(items, its...)
	}
	return items, nil
}
