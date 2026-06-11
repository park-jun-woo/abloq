//ff:func feature=quest type=rule control=sequence
//ff:what firstFact 검증 — 단건 그대로, 다건은 첫 Fact Actual에 "(외 N건)" 병기
package translation

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestFirstFact(t *testing.T) {
	one := []quest.Fact{{Where: "w", Actual: "a"}}
	if f := firstFact(one); f.Actual != "a" {
		t.Errorf("single = %+v", f)
	}
	three := []quest.Fact{{Actual: "a"}, {Actual: "b"}, {Actual: "c"}}
	if f := firstFact(three); f.Actual != "a (외 2건)" {
		t.Errorf("multi = %+v", f)
	}
}
