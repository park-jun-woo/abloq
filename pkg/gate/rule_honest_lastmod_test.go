//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what [honest-lastmod] 베이스라인 픽스처 케이스 — lastmod 불변 PASS, 실변경 동반 PASS, 사소-diff 갱신 FAIL 검증
package gate

import "testing"

func TestRuleHonestLastmod(t *testing.T) {
	cases := []struct {
		name, newFile string
		wantDiags     int
		wantMsgPart   string
	}{
		{"pass unchanged lastmod", "baseline/reordered.md", 0, ""},
		{"pass meaningful update", "baseline/lastmod-real.md", 0, ""},
		{"fail lastmod only", "baseline/lastmod-only.md", 1, "min_meaningful_diff"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkBaselineRule(t, ruleHonestLastmod, tc.newFile, "honest-lastmod", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
