//ff:func feature=quest type=rule control=sequence
//ff:what parityFacts 검증 — 클린 쌍 0건, 복합 훼손(코드 변조+lastmod 위조)에서 ①~⑦ 순서로 다건 수집
package translation

import (
	"strings"
	"testing"
)

func TestParityFacts(t *testing.T) {
	root := writeInstance(t)
	origin, ko := passPair()
	if facts := parityFacts(subWith(t, root, origin, ko)); len(facts) != 0 {
		t.Errorf("clean pair: %+v", facts)
	}
	broken := strings.Replace(ko, `echo "do not translate"`, "echo translated", 1)
	broken = strings.Replace(broken, "lastmod: 2026-06-03", "lastmod: 2026-06-09", 1)
	facts := parityFacts(subWith(t, root, origin, broken))
	if len(facts) != 3 {
		t.Fatalf("facts = %d, want 3 (code missing+injected, lastmod)", len(facts))
	}
	if !strings.Contains(facts[0].Where, "#code") || !strings.Contains(facts[2].Where, "#lastmod") {
		t.Errorf("order: %+v", facts)
	}
}
