//ff:func feature=quest type=rule control=sequence
//ff:what image-first 어댑터 발동 검증 — 첫 본문 라인이 이미지가 아닌 글에서 Fact 매핑
package writing

import "testing"

func TestAdaptImageFirst(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	noImage := removeLine(art, "![main](cover.png)")
	fired, fact := fireRule(t, adaptRule("image-first"), subWith(t, root, noImage, ""))
	if !fired {
		t.Fatal("image-first: want fired without a main image")
	}
	if fact.Expected == "" || fact.Actual == "" || fact.Where == "" {
		t.Errorf("Fact incomplete: %+v", fact)
	}
}
