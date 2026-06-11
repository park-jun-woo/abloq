//ff:func feature=quest type=rule control=sequence
//ff:what checkFMMirror 검증 — date 불일치·lastmod 부재가 각각 Fact, 원문 이식 쌍은 무발동
package translation

import (
	"strings"
	"testing"
)

func TestCheckFMMirror(t *testing.T) {
	root := writeInstance(t)
	origin, ko := passPair()
	if facts := checkFMMirror(subWith(t, root, origin, ko)); len(facts) != 0 {
		t.Errorf("mirrored pair: %+v", facts)
	}
	dateDrift := strings.Replace(ko, "date: 2026-06-01", "date: 2026-06-02", 1)
	facts := checkFMMirror(subWith(t, root, origin, dateDrift))
	if len(facts) != 1 || !strings.Contains(facts[0].Where, "#date") {
		t.Errorf("date drift: %+v", facts)
	}
	noLastmod := removeLine(ko, "lastmod:")
	facts = checkFMMirror(subWith(t, root, origin, noLastmod))
	if len(facts) != 1 || !strings.Contains(facts[0].Actual, "missing or unparseable") {
		t.Errorf("lastmod missing: %+v", facts)
	}
}
