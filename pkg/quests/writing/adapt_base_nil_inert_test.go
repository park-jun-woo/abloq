//ff:func feature=quest type=rule control=iteration dimension=1
//ff:what Base nil 규약 검증 — 기준선 3룰(section-preserved·body-lossless·front-matter-intact)은 신규 글에서 inert
package writing

import "testing"

func TestAdaptBaseNilInert(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	bare := removeLine(art, "## Changelog") // baseline rules must stay silent anyway
	for _, id := range []string{"section-preserved", "body-lossless", "front-matter-intact"} {
		if fired, _ := fireRule(t, adaptRule(id), subWith(t, root, bare, "")); fired {
			t.Errorf("%s: fired with Base nil (must be inert on new articles)", id)
		}
	}
}
