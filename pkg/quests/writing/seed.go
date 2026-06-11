//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what Seed — insight.yaml 경로 인자마다 Item 1개(Key=lang/section/slug) 시드, 인자 없음·명세 불량은 에러
package writing

import (
	"fmt"

	"github.com/park-jun-woo/reins/pkg/quest"
)

// Seed creates one TODO item per insight.yaml path argument. The spec is
// loaded and validated up front so a broken spec never enters the ratchet.
func (Definition) Seed(args []string) ([]*quest.Item, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("usage: scan <insight.yaml> [insight.yaml...]")
	}
	var items []*quest.Item
	for _, arg := range args {
		it, err := seedItem(arg)
		if err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}
