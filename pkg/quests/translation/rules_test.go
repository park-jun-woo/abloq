//ff:func feature=quest type=frame control=iteration dimension=1
//ff:what Rules 카탈로그 검증 — parity·스코프드 slug·채택 6룰·hugo-build 순서 고정, 전부 LevelFail, 배제 룰 부재
package translation

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRules(t *testing.T) {
	rules := Definition{}.Rules()
	want := []string{
		"translation-parity", "slug-consistency", "front-matter-schema",
		"image-first", "image-attribution", "section-order",
		"heading-canonical", "min-sources", "hugo-build",
	}
	if len(rules) != len(want) {
		t.Fatalf("rules = %d, want %d", len(rules), len(want))
	}
	excluded := map[string]bool{
		"section-preserved": true, "body-lossless": true, "front-matter-intact": true,
		"honest-lastmod": true, "hreflang-complete": true,
		"numeric-claim-sourced": true, "citation-exists": true,
	}
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
