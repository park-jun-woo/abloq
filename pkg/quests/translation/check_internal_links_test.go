//ff:func feature=quest type=rule control=sequence
//ff:what checkInternalLinks 검증 — 프리픽스 누락 Fact(⑥a)와 정규화 후 multiset 불일치 Fact(⑥b), 클린 쌍 무발동
package translation

import (
	"strings"
	"testing"
)

func TestCheckInternalLinks(t *testing.T) {
	root := writeInstance(t)
	origin, ko := passPair()
	if facts := checkInternalLinks(subWith(t, root, origin, ko)); len(facts) != 0 {
		t.Errorf("clean pair: %+v", facts)
	}
	unprefixed := strings.Replace(ko, "(/ko/posts/first-post/)", "(/posts/first-post/)", 1)
	facts := checkInternalLinks(subWith(t, root, origin, unprefixed))
	if len(facts) != 1 || !strings.Contains(facts[0].Actual, "unprefixed link") {
		t.Errorf("unprefixed: %+v", facts)
	}
	retargeted := strings.Replace(ko, "(/ko/posts/first-post/)", "(/ko/posts/other-post/)", 1)
	facts = checkInternalLinks(subWith(t, root, origin, retargeted))
	if len(facts) != 2 || !strings.Contains(facts[0].Actual, "missing in translation") {
		t.Errorf("retargeted: %+v", facts)
	}
}
