//ff:func feature=quest type=rule control=sequence
//ff:what image-attribution 어댑터 발동 검증 — 메인 이미지 직후 이탤릭 표기 없는 글에서 Fact 매핑
package writing

import (
	"strings"
	"testing"
)

func TestAdaptImageAttribution(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	noAttrib := removeLine(art, "*Image: by Tester*")
	fired, fact := fireRule(t, adaptRule("image-attribution"), subWith(t, root, noAttrib, ""))
	if !fired {
		t.Fatal("image-attribution: want fired without an attribution line")
	}
	if !strings.Contains(fact.Actual, "attribution") {
		t.Errorf("Actual = %q", fact.Actual)
	}
}
