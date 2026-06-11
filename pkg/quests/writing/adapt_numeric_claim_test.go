//ff:func feature=quest type=rule control=sequence
//ff:what numeric-claim-sourced 어댑터 발동 검증 — 무출처 수치 주장(40 percent improved) 문단에서 Fact 매핑
package writing

import (
	"strings"
	"testing"
)

func TestAdaptNumericClaim(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	claim := strings.Replace(art,
		"This body mentions the alpha anchor.",
		"This body mentions the alpha anchor.\n\nThroughput improved by 40 percent after the change.",
		1)
	fired, fact := fireRule(t, adaptRule("numeric-claim-sourced"), subWith(t, root, claim, ""))
	if !fired {
		t.Fatal("numeric-claim-sourced: want fired on an unsourced numeric claim")
	}
	if fact.Where == "" || fact.Actual == "" {
		t.Errorf("Fact incomplete: %+v", fact)
	}
}
