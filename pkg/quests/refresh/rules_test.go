//ff:func feature=quest type=frame control=iteration dimension=1 topic=queue
//ff:what Rules 카탈로그 검증 — 채택 14룰의 ID·순서·전부 LevelFail, body-lossless 비편입
package refresh

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRules(t *testing.T) {
	rules := Definition{}.Rules()
	want := []string{
		"lastmod-advance", "claim-preserved", "queue-scope",
		"honest-lastmod", "front-matter-intact", "front-matter-schema",
		"section-order", "section-preserved", "heading-canonical",
		"image-first", "image-attribution", "min-sources",
		"numeric-claim-sourced", "citation-exists",
	}
	if len(rules) != len(want) {
		t.Fatalf("rule count = %d, want %d", len(rules), len(want))
	}
	for i, r := range rules {
		if r.Meta.ID != want[i] {
			t.Errorf("rules[%d] = %s, want %s", i, r.Meta.ID, want[i])
		}
		if r.Meta.Level != rgate.LevelFail {
			t.Errorf("%s: level must be Fail (Review deadlocks the ratchet)", r.Meta.ID)
		}
		if r.Meta.ID == "body-lossless" {
			t.Error("body-lossless must be excluded — refreshing rewrites lines in place")
		}
	}
}
