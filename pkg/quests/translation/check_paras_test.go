//ff:func feature=quest type=rule control=sequence
//ff:what checkParas 검증 — 문단 블록 수 불일치 Fact(기대/실측 수치 포함), 일치 쌍은 무발동
package translation

import (
	"strings"
	"testing"
)

func TestCheckParas(t *testing.T) {
	origin, ko := passPair()
	o, k := docOf(t, "en", origin), docOf(t, "ko", ko)
	if facts := checkParas("w", o, k); len(facts) != 0 {
		t.Errorf("clean pair: %+v", facts)
	}
	merged := docOf(t, "ko", strings.Replace(ko, "링크가 있는 도입 문단.\n\n## 준비", "링크가 있는 도입 문단.\n## 준비", 1))
	facts := checkParas("w", o, merged)
	if len(facts) != 1 || !strings.Contains(facts[0].Where, "#paragraphs") {
		t.Errorf("merged blocks: %+v", facts)
	}
}
