//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what [front-matter-intact] 베이스라인 픽스처 케이스 — lastmod만 변경 PASS 2건, title 변경 FAIL 검증
package gate

import "testing"

func TestRuleFrontMatterIntact(t *testing.T) {
	cases := []struct {
		name, newFile string
		wantDiags     int
		wantMsgPart   string
	}{
		{"pass identical", "baseline/base.md", 0, ""},
		{"pass lastmod only", "baseline/lastmod-only.md", 0, ""},
		{"fail title changed", "baseline/fm-changed.md", 1, "only lastmod may change"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkBaselineRule(t, ruleFrontMatterIntact, tc.newFile, "front-matter-intact", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
