//ff:func feature=quest type=rule control=sequence
//ff:what translation-parity 검출 ⑦ — 번역 lastmod 위조(원문과 불일치) 시 발동, Fact Where가 #lastmod인지
package translation

import (
	"strings"
	"testing"
)

func TestParityLastmodForged(t *testing.T) {
	origin, ko := passPair()
	broken := strings.Replace(ko, "lastmod: 2026-06-03", "lastmod: 2026-06-09", 1)
	fired, fact := fireRule(t, ruleTranslationParity(), subWith(t, writeInstance(t), origin, broken))
	if !fired {
		t.Fatal("forged lastmod: want translation-parity fired")
	}
	if !strings.Contains(fact.Where, "#lastmod") {
		t.Errorf("Where = %q, want #lastmod", fact.Where)
	}
	if !strings.Contains(fact.Expected, "2026-06-03") {
		t.Errorf("Expected = %q, want the origin lastmod quoted", fact.Expected)
	}
}
