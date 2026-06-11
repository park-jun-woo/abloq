//ff:func feature=quest type=frame control=iteration dimension=1
//ff:what Rules 카탈로그 검증 — 채택 11룰+review-record 순서 고정, 전부 LevelFail, 배제 3룰(slug-consistency·hreflang-complete·honest-lastmod) 부재
package writing

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRules(t *testing.T) {
	rules := Definition{}.Rules()
	want := []string{
		"image-first", "image-attribution", "section-order", "section-preserved",
		"body-lossless", "front-matter-intact", "heading-canonical",
		"front-matter-schema", "min-sources", "numeric-claim-sourced",
		"citation-exists", "review-record",
	}
	if len(rules) != len(want) {
		t.Fatalf("rules = %d, want %d", len(rules), len(want))
	}
	excluded := map[string]bool{"slug-consistency": true, "hreflang-complete": true, "honest-lastmod": true}
	for i, r := range rules {
		if r.Meta.ID != want[i] {
			t.Errorf("rules[%d] = %s, want %s", i, r.Meta.ID, want[i])
		}
		if r.Meta.Level != rgate.LevelFail {
			t.Errorf("%s: Level = %v, want LevelFail", r.Meta.ID, r.Meta.Level)
		}
		if excluded[r.Meta.ID] {
			t.Errorf("excluded rule %s present in catalog", r.Meta.ID)
		}
		if r.Meta.Desc == "" {
			t.Errorf("%s: empty Desc", r.Meta.ID)
		}
	}
}
