//ff:func feature=quest type=rule control=sequence
//ff:what translation-parity 검출 ① — 번역에서 sources 섹션(헤딩) 누락 시 발동, Fact가 대상 글에 위치하는지
package translation

import (
	"strings"
	"testing"
)

func TestParitySectionMissing(t *testing.T) {
	origin, ko := passPair()
	broken := removeLine(ko, "## 출처")
	fired, fact := fireRule(t, ruleTranslationParity(), subWith(t, writeInstance(t), origin, broken))
	if !fired {
		t.Fatal("section dropped: want translation-parity fired")
	}
	if !strings.HasPrefix(fact.Where, "content/ko/posts/fixture.md#") {
		t.Errorf("Where = %q, want located at the translation", fact.Where)
	}
	if fact.Expected == "" || fact.Actual == "" {
		t.Errorf("Fact incomplete: %+v", fact)
	}
}
