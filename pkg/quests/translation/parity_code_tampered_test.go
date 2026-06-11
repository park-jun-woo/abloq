//ff:func feature=quest type=rule control=sequence
//ff:what translation-parity 검출 ④ — 코드블록 내용 변조(번역) 시 발동, Fact Where가 #code인지
package translation

import (
	"strings"
	"testing"
)

func TestParityCodeTampered(t *testing.T) {
	origin, ko := passPair()
	broken := strings.Replace(ko, `echo "do not translate"`, `echo "번역해 버림"`, 1)
	fired, fact := fireRule(t, ruleTranslationParity(), subWith(t, writeInstance(t), origin, broken))
	if !fired {
		t.Fatal("code block tampered: want translation-parity fired")
	}
	if !strings.Contains(fact.Where, "#code") {
		t.Errorf("Where = %q, want #code", fact.Where)
	}
}
