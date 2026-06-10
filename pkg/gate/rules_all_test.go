//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what Rules 레지스트리가 11룰을 고정 순서·중복 없는 ID로 노출하는지 검증
package gate

import "testing"

func TestRulesAll(t *testing.T) {
	rules := Rules()
	want := []string{
		"image-first", "image-attribution", "section-order", "section-preserved",
		"body-lossless", "front-matter-intact", "heading-canonical",
		"front-matter-schema", "slug-consistency", "honest-lastmod", "hreflang-complete",
	}
	if len(rules) != len(want) {
		t.Fatalf("want %d rules, got %d", len(want), len(rules))
	}
	for i, r := range rules {
		if r.ID != want[i] {
			t.Errorf("rules[%d].ID = %s, want %s", i, r.ID, want[i])
		}
		if r.Desc == "" || r.Check == nil {
			t.Errorf("rule %s lacks desc or check", r.ID)
		}
	}
}
