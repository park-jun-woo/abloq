//ff:func feature=quest type=rule control=sequence
//ff:what translation-parity 검출 ⑥(a) — 내부 글 링크 미치환(자기 언어 프리픽스 누락) 시 발동, unprefixed Fact인지
package translation

import (
	"strings"
	"testing"
)

func TestParityUnprefixedLink(t *testing.T) {
	origin, ko := passPair()
	broken := strings.Replace(ko, "(/ko/posts/first-post/)", "(/posts/first-post/)", 1)
	fired, fact := fireRule(t, ruleTranslationParity(), subWith(t, writeInstance(t), origin, broken))
	if !fired {
		t.Fatal("unprefixed internal link: want translation-parity fired")
	}
	if !strings.Contains(fact.Actual, "unprefixed link: /posts/first-post/") {
		t.Errorf("Actual = %q, want the unprefixed link named", fact.Actual)
	}
}
