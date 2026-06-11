//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Seed — quests/queue/의 kind=refresh 큐 파일마다 Item 1개(공통 SeedQueue), priority 내림차순
package refresh

import (
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Seed scans the instance's quests/queue/ directory and seeds one TODO item
// per kind=refresh queue file, highest priority first. The queue payload is
// frozen into the item at seed time (shared consumption protocol).
func (Definition) Seed(args []string) ([]*quest.Item, error) {
	return common.SeedQueue("refresh", args)
}
