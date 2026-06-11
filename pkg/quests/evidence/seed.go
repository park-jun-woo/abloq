//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Seed — quests/queue/의 kind=evidence 큐 파일마다 Item 1개(공통 SeedQueue), priority 내림차순
package evidence

import (
	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// Seed scans the instance's quests/queue/ directory and seeds one TODO item
// per kind=evidence queue file, highest priority first. The queue payload
// (claims hashes + rot URLs) is frozen into the item at seed time (shared
// consumption protocol).
func (Definition) Seed(args []string) ([]*quest.Item, error) {
	return common.SeedQueue("evidence", args)
}
