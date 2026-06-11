//ff:func feature=quest type=rule control=sequence
//ff:what checkHeadings 검증 — 레벨 시퀀스 불일치와 인식 섹션 키 불일치를 각각 Fact로, 일치 쌍은 무발동
package translation

import (
	"strings"
	"testing"
)

func TestCheckHeadings(t *testing.T) {
	origin, ko := passPair()
	o, k := docOf(t, "en", origin), docOf(t, "ko", ko)
	if facts := checkHeadings("w", o, k); len(facts) != 0 {
		t.Errorf("clean pair: %+v", facts)
	}
	demoted := docOf(t, "ko", strings.Replace(ko, "## 준비", "### 준비", 1))
	facts := checkHeadings("w", o, demoted)
	if len(facts) != 1 || !strings.Contains(facts[0].Where, "#headings") {
		t.Errorf("level drift: %+v", facts)
	}
	wrongHead := docOf(t, "ko", strings.Replace(ko, "## 출처", "## 출처들", 1))
	facts = checkHeadings("w", o, wrongHead)
	if len(facts) != 1 || !strings.Contains(facts[0].Where, "#sections") {
		t.Errorf("unrecognized sources heading: %+v", facts)
	}
}
