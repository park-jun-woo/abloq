//ff:func feature=quest type=rule control=sequence
//ff:what checkMultiset 검증 — 누락(원문→번역)과 주입(번역→원문)이 각각 Fact, 동일 multiset은 무발동
package translation

import (
	"strings"
	"testing"
)

func TestCheckMultiset(t *testing.T) {
	if facts := checkMultiset("w", "x", []string{"a", "a"}, []string{"a", "a"}); len(facts) != 0 {
		t.Errorf("equal multisets: %+v", facts)
	}
	facts := checkMultiset("w", "x", []string{"a", "b"}, []string{"a", "c"})
	if len(facts) != 2 {
		t.Fatalf("facts = %d, want 2 (missing + injected)", len(facts))
	}
	if !strings.Contains(facts[0].Actual, "missing in translation: b") {
		t.Errorf("missing fact: %+v", facts[0])
	}
	if !strings.Contains(facts[1].Actual, "injected in translation: c") {
		t.Errorf("injected fact: %+v", facts[1])
	}
}
