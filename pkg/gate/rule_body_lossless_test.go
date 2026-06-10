//ff:func feature=gate type=rule control=iteration dimension=1 topic=lossless
//ff:what [body-lossless] 베이스라인 픽스처 케이스 — 재배열·추가 PASS, 본문 라인 삭제 FAIL 검증
package gate

import "testing"

func TestRuleBodyLossless(t *testing.T) {
	cases := []struct {
		name, newFile string
		wantDiags     int
		wantMsgPart   string
	}{
		{"pass reordered", "baseline/reordered.md", 0, ""},
		{"pass dropped heading keeps body", "baseline/dropped-section.md", 0, ""},
		{"pass added paragraph", "baseline/lastmod-real.md", 0, ""},
		{"fail deleted line", "baseline/deleted-line.md", 1, "deleted or altered"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkBaselineRule(t, ruleBodyLossless, tc.newFile, "body-lossless", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
