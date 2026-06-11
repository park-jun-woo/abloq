//ff:func feature=quest type=rule control=sequence
//ff:what body-lossless 어댑터 발동 검증 — Base 부착 픽스처에서 기준선 본문 라인 삭제 시 Fact 매핑
package writing

import "testing"

func TestAdaptBodyLossless(t *testing.T) {
	root := writeInstance(t)
	base, _ := passFixtures()
	cur := removeLine(base, "This body mentions the alpha anchor")
	fired, fact := fireRule(t, adaptRule("body-lossless"), subWith(t, root, cur, base))
	if !fired {
		t.Fatal("body-lossless: want fired when a baseline body line is deleted")
	}
	if fact.Where == "" || fact.Actual == "" {
		t.Errorf("Fact incomplete: %+v", fact)
	}
}
