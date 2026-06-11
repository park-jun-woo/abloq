//ff:func feature=quest type=rule control=sequence
//ff:what translation-parity 검출 ⑤ — 외부 링크 URL 변조 시 발동(누락+주입 양방향), Fact Where가 #external-links인지
package translation

import (
	"strings"
	"testing"
)

func TestParityExternalLink(t *testing.T) {
	origin, ko := passPair()
	broken := strings.Replace(ko, "(https://example.org/spec) 링크", "(https://evil.example/spec) 링크", 1)
	fired, fact := fireRule(t, ruleTranslationParity(), subWith(t, writeInstance(t), origin, broken))
	if !fired {
		t.Fatal("external URL tampered: want translation-parity fired")
	}
	if !strings.Contains(fact.Where, "#external-links") {
		t.Errorf("Where = %q, want #external-links", fact.Where)
	}
	if !strings.Contains(fact.Actual, "(외 1건)") {
		t.Errorf("Actual = %q, want the injected counterpart counted", fact.Actual)
	}
}
