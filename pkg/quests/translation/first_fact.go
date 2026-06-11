//ff:func feature=quest type=rule control=sequence
//ff:what Fact 목록 → 룰당 Fact 1건 규약으로 축약 — 첫 Fact에 "(외 N건)" 병기
package translation

import (
	"fmt"

	"github.com/park-jun-woo/reins/pkg/quest"
)

// firstFact collapses a violation list to the adapter convention's single
// Fact: the first one, with the remaining count noted.
func firstFact(facts []quest.Fact) quest.Fact {
	f := facts[0]
	if len(facts) > 1 {
		f.Actual += fmt.Sprintf(" (외 %d건)", len(facts)-1)
	}
	return f
}
